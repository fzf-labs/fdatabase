package batch

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// MysqlBatchUpdateToSQLArray 批量更新
// tableName: 表名
// dataList: 需要更新的数据列表,必须是指向结构体的切片
func MysqlBatchUpdateToSQLArray(tableName string, dataList any) ([]string, error) {
	if tableName == "" {
		return nil, errors.New("tableName cannot be empty")
	}

	// 检查 dataList 是否为切片
	rv := reflect.ValueOf(dataList)
	if rv.Kind() != reflect.Slice {
		return nil, errors.New("dataList must be a slice")
	}

	if rv.Len() == 0 {
		return nil, errors.New("dataList cannot be empty")
	}

	// 获取元素类型
	elemType := rv.Type().Elem()
	if elemType.Kind() != reflect.Ptr || elemType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("dataList must be a slice of struct pointers")
	}

	// 获取结构体字段信息
	fields, err := getStructFields(elemType.Elem())
	if err != nil {
		return nil, fmt.Errorf("get struct fields error: %w", err)
	}

	// 检查是否存在 "id" 字段
	if _, ok := fields["id"]; !ok {
		return nil, errors.New("struct must have a field with json tag 'id'")
	}

	// 准备数据
	ids := make([]string, 0, rv.Len())
	updateMap := make(map[string][]string)
	for i := 0; i < rv.Len(); i++ {
		// 获取每个结构体实例
		structVal := rv.Index(i).Elem()
		idField := structVal.FieldByName(fields["id"].name)
		if !idField.IsValid() {
			return nil, fmt.Errorf("id field not found in struct at index %d", i)
		}

		id, err := strconv.ParseInt(idField.String(), 10, 64)
		if err != nil || id < 1 {
			return nil, fmt.Errorf("invalid id value at index %d", i)
		}
		ids = append(ids, strconv.FormatInt(id, 10))

		// 处理其他字段
		for fieldName, fieldInfo := range fields {
			if fieldName == "id" {
				continue
			}
			fieldValue := structVal.FieldByName(fieldInfo.name)
			if !fieldValue.IsValid() {
				return nil, fmt.Errorf("field %s not found in struct at index %d", fieldName, i)
			}

			valStr, err := formatFieldValue(fieldValue)
			if err != nil {
				return nil, fmt.Errorf("format field %s error at index %d: %w", fieldName, i, err)
			}
			updateMap[fieldName] = append(updateMap[fieldName], valStr)
		}
	}

	// 计算 SQL 语句数量
	length := len(ids)
	const batchSize = 200
	sqlQuantity := getSQLQuantity(length, batchSize)

	// 生成 SQL 语句
	sqlArray := make([]string, 0, sqlQuantity)
	for i := 0; i < sqlQuantity; i++ {
		batchStart := i * batchSize
		batchEnd := min((i+1)*batchSize, length)

		sql, err := buildBatchUpdateSQL(tableName, updateMap, ids[batchStart:batchEnd])
		if err != nil {
			return nil, fmt.Errorf("build batch update SQL error: %w", err)
		}
		sqlArray = append(sqlArray, sql)
	}

	return sqlArray, nil
}

// 辅助函数：获取结构体字段信息
func getStructFields(structType reflect.Type) (map[string]structField, error) {
	fields := make(map[string]structField)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			return nil, fmt.Errorf("field %s must have a json tag", field.Name)
		}
		if _, ok := fields[jsonTag]; ok {
			return nil, fmt.Errorf("duplicate json tag '%s' in struct", jsonTag)
		}
		fields[jsonTag] = structField{
			name:     field.Name,
			typeKind: field.Type.Kind(),
		}
	}
	return fields, nil
}

type structField struct {
	name     string
	typeKind reflect.Kind
}

// 辅助函数：格式化字段值
func formatFieldValue(field reflect.Value) (string, error) {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10), nil
	case reflect.String:
		return fmt.Sprintf("'%s'", strings.ReplaceAll(field.String(), "'", "\\'")), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(field.Bool()), nil
	default:
		return "", fmt.Errorf("unsupported field type: %v", field.Kind())
	}
}

// 辅助函数：计算 SQL 语句数量
func getSQLQuantity(length, batchSize int) int {
	return int(math.Ceil(float64(length) / float64(batchSize)))
}

// 辅助函数：生成每批次的 SQL 语句
func buildBatchUpdateSQL(tableName string, updateMap map[string][]string, batchIDs []string) (string, error) {
	if len(batchIDs) == 0 {
		return "", errors.New("batchIDs cannot be empty")
	}

	var sqlBuilder strings.Builder
	sqlBuilder.Grow(4096) // 预分配空间

	sqlBuilder.WriteString("UPDATE " + escapeIdentifier(tableName) + " SET ")

	setClauses := make([]string, 0, len(updateMap))
	for fieldName, fieldValueList := range updateMap {
		if len(fieldValueList) != len(batchIDs) {
			return "", fmt.Errorf("field %s values count does not match IDs count", fieldName)
		}

		clause := escapeIdentifier(fieldName) + " = CASE id"
		for i, id := range batchIDs {
			clause += " WHEN " + id + " THEN " + fieldValueList[i]
		}
		clause += " END"
		setClauses = append(setClauses, clause)
	}

	sqlBuilder.WriteString(strings.Join(setClauses, ", "))
	sqlBuilder.WriteString(" WHERE id IN (" + strings.Join(batchIDs, ",") + ")")

	return sqlBuilder.String(), nil
}

func escapeIdentifier(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
