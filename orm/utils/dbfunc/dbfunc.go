package dbfunc

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	Postgres = "postgres"
	MySQL    = "mysql"
)

type Index struct {
	TableName  string `json:"table_name" gorm:"column:table_name"`
	IndexName  string `json:"index_name" gorm:"column:index_name"`
	ColumnName string `json:"column_name" gorm:"column:column_name"`
	IsUnique   bool   `json:"is_unique" gorm:"column:is_unique"`
	Primary    bool   `json:"primary" gorm:"column:primary"`
}

// GetIndexes 获取索引
func GetIndexes(db *gorm.DB, table string) ([]*Index, error) {
	resp := make([]*Index, 0)
	var err error
	switch db.Dialector.Name() {
	case Postgres:
		resp, err = GetPgIndexes(db, table)
		if err != nil {
			return nil, err
		}
	default:
		result, err := db.Migrator().GetIndexes(table)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			unique, _ := v.Unique()
			isPrimaryKey, _ := v.PrimaryKey()
			for _, vv := range v.Columns() {
				resp = append(resp, &Index{
					TableName:  table,
					IndexName:  v.Name(),
					ColumnName: vv,
					IsUnique:   unique,
					Primary:    isPrimaryKey,
				})
			}
		}
	}
	return resp, nil
}

// GetPgIndexes 查询PG索引
func GetPgIndexes(db *gorm.DB, table string) ([]*Index, error) {
	result := make([]*Index, 0)
	sql := fmt.Sprintf(`select t.relname as table_name,i.relname as index_name,a.attname as column_name,ix.indisunique as is_unique,ix.indisprimary as primary from pg_class t,pg_class i,pg_index ix,pg_attribute a where t.oid=ix.indrelid and i.oid=ix.indexrelid and a.attrelid=t.oid and a.attnum=ANY(ix.indkey)and t.relkind='r' and t.relname='%s'`, table)
	err := db.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SortIndexColumns 排序索引字段
func SortIndexColumns(db *gorm.DB, table string) (map[string][]string, error) {
	resp := make(map[string][]string)
	var err error
	switch db.Dialector.Name() {
	case Postgres:
		resp, err = pgSortIndexColumns(db, table)
		if err != nil {
			return nil, err
		}
	default:
		result, err := db.Migrator().GetIndexes(table)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			if _, ok := resp[v.Name()]; !ok {
				resp[v.Name()] = make([]string, 0)
			}
			resp[v.Name()] = v.Columns()
		}
	}
	return resp, nil
}

// pgSortIndexColumns  postgres索引字段排序
func pgSortIndexColumns(db *gorm.DB, table string) (map[string][]string, error) {
	resp := make(map[string][]string)
	type Tmp struct {
		TableName  string `gorm:"column:table_name" json:"table_name"`
		IndexName  string `gorm:"column:index_name" json:"index_name"`
		ColumnName string `gorm:"column:column_name" json:"column_name"`
	}
	result := make([]Tmp, 0)
	sql := fmt.Sprintf(`SELECT t.relname AS table_name,i.relname AS index_name,a.attname AS column_name,ix.indisunique AS is_unique,ix.indisprimary AS PRIMARY FROM pg_class t JOIN pg_index ix ON t.oid=ix.indrelid JOIN pg_class i ON i.oid=ix.indexrelid JOIN pg_attribute a ON a.attrelid=t.oid AND a.attnum=ANY(ix.indkey)WHERE t.relkind='r' AND t.relname='%s' ORDER BY ix.indrelid,(array_position(ix.indkey,a.attnum))`, table)
	err := db.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if _, ok := resp[v.IndexName]; !ok {
			resp[v.IndexName] = make([]string, 0)
		}
		resp[v.IndexName] = append(resp[v.IndexName], v.ColumnName)
	}
	return resp, nil
}

// GetPartitionChildTableForTable 获取PG分区表的子表
func GetPartitionChildTableForTable(db *gorm.DB, tableName string) ([]string, error) {
	switch db.Dialector.Name() {
	case Postgres:
		resp, err := getPGPartitionChildTableForTable(db, tableName)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case MySQL:
		resp, err := getMySQLPartitionChildTableForTable(db, tableName)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, nil
	}
}

// getPGPartitionChildTableForTable 获取PG分区表的子表
func getPGPartitionChildTableForTable(db *gorm.DB, tableName string) ([]string, error) {
	result := make([]string, 0)
	sql := fmt.Sprintf(`SELECT c.relname AS child_table FROM pg_catalog.pg_class c JOIN pg_catalog.pg_inherits i ON c.oid=i.inhrelid WHERE i.inhparent=(SELECT oid FROM pg_class WHERE relname='%s')`, tableName)
	err := db.Raw(sql).Pluck("child_table", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMySQLPartitionChildTableForTable 获取MySQL分区表的子表
func getMySQLPartitionChildTableForTable(db *gorm.DB, tableName string) ([]string, error) {
	result := make([]string, 0)
	sql := fmt.Sprintf(`SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_TYPE='BASE TABLE' AND TABLE_SCHEMA='%s' AND TABLE_NAME LIKE '%s'`, db.Migrator().CurrentDatabase(), tableName+"%")
	err := db.Raw(sql).Pluck("TABLE_NAME", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPartitionChildTable 获取所有分区表的子表
func GetPartitionChildTable(db *gorm.DB) (resp []string, err error) {
	switch db.Dialector.Name() {
	case Postgres:
		resp, err = getPGPartitionChildTable(db)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case MySQL:
		resp, err = getMySQLPartitionChildTable(db)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, nil
	}
}

// getPGPartitionChildTable 获取PG获取所有分区表的子表
func getPGPartitionChildTable(db *gorm.DB) ([]string, error) {
	result := make([]string, 0)
	sql := `SELECT c.relname AS child_table FROM pg_catalog.pg_class c JOIN pg_catalog.pg_inherits i ON c.oid=i.inhrelid`
	err := db.Raw(sql).Pluck("child_table", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getMySQLPartitionChildTable 获取MySQL获取所有分区表的子表
func getMySQLPartitionChildTable(db *gorm.DB) ([]string, error) {
	result := make([]string, 0)
	sql := `SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_TYPE='BASE TABLE' AND TABLE_SCHEMA='%s'`
	err := db.Raw(sql).Pluck("TABLE_NAME", &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPartitionTableToChildTables 获取分区表到子表的映射
func GetPartitionTableToChildTables(db *gorm.DB) (resp map[string][]string, err error) {
	switch db.Dialector.Name() {
	case Postgres:
		resp, err = getPGPartitionTableToChildTables(db)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case MySQL:
		resp, err = getMySQLPartitionTableToChildTables(db)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, nil
	}
}

// getPartitionTableToChildTable 获取PG分区表到子表的映射
func getPGPartitionTableToChildTables(db *gorm.DB) (map[string][]string, error) {
	resp := make(map[string][]string)
	type tmp struct {
		PartitionedTable string `gorm:"column:partitioned_table" json:"partitioned_table"`
		ChildTables      string `gorm:"column:child_tables" json:"child_tables"`
	}
	result := make([]tmp, 0)
	sql := `SELECT p.relname AS partitioned_table,array_to_string(array_agg(c.relname),',')AS child_tables FROM pg_catalog.pg_class c JOIN pg_catalog.pg_inherits i ON c.oid=i.inhrelid JOIN pg_catalog.pg_class p ON p.oid=i.inhparent GROUP BY p.relname;`
	err := db.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		resp[v.PartitionedTable] = append(resp[v.PartitionedTable], strings.Split(v.ChildTables, ",")...)
	}
	return resp, nil
}

// getMySQLPartitionTableToChildTable 获取MySQL分区表到子表的映射
func getMySQLPartitionTableToChildTables(db *gorm.DB) (map[string][]string, error) {
	resp := make(map[string][]string)
	type tmp struct {
		PartitionedTable string `gorm:"column:partitioned_table" json:"partitioned_table"`
		ChildTables      string `gorm:"column:child_tables" json:"child_tables"`
	}
	result := make([]tmp, 0)
	sql := `SELECT TABLE_NAME AS partitioned_table,GROUP_CONCAT(TABLE_NAME) AS child_tables FROM information_schema.TABLES WHERE TABLE_TYPE='BASE TABLE' AND TABLE_SCHEMA='%s' GROUP BY TABLE_NAME;`
	err := db.Raw(sql).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		resp[v.PartitionedTable] = append(resp[v.PartitionedTable], strings.Split(v.ChildTables, ",")...)
	}
	return resp, nil
}
