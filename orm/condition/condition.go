package condition

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"gorm.io/gorm/clause"
)

type Exp string // 操作

const (
	EQ      Exp = "="
	NEQ     Exp = "!="
	GT      Exp = ">"
	GTE     Exp = ">="
	LT      Exp = "<"
	LTE     Exp = "<="
	IN      Exp = "IN"
	NOTIN   Exp = "NOT IN"
	LIKE    Exp = "LIKE"
	NOTLIKE Exp = "NOT LIKE"
)

type Logic string // 逻辑关系

const (
	AND Logic = "AND" // 逻辑关系 and
	OR  Logic = "OR"  // 逻辑关系 or
)

type Order string // 排序

const (
	ASC  Order = "ASC"  // 升序
	DESC Order = "DESC" // 降序
)

// QueryParam 查询条件
type QueryParam struct {
	Field string `json:"field"` // 字段
	Value any    `json:"value"` // 值（当Exp为IN, NOTIN 时为[]any）
	Exp   Exp    `json:"exp"`   // 操作 "=", "!=", ">", ">=", "<", "<=", "IN", "NOT IN", "LIKE", "NOT LIKE"
	Logic Logic  `json:"logic"` // 逻辑关系 AND OR
}

// OrderParam 排序条件
type OrderParam struct {
	Field string `json:"field"` // 字段
	Order Order  `json:"order"` // 排序 ASC DESC
}

// Req 请求-自定义查询
type Req struct {
	Page     int32         `json:"page"`     // 页码
	PageSize int32         `json:"pageSize"` // 页数
	Query    []*QueryParam `json:"query"`    // 查询条件
	Order    []*OrderParam `json:"order"`    // 排序条件
}

// Reply 返回-自定义查询
type Reply struct {
	Page      int32 `json:"page"`      // 第几页
	PageSize  int32 `json:"pageSize"`  // 页大小
	Total     int32 `json:"total"`     // 总数
	PrevPage  int32 `json:"prevPage"`  // 上一页
	NextPage  int32 `json:"nextPage"`  // 下一页
	TotalPage int32 `json:"totalPage"` // 总页数
}

// ExpValidate 验证Exp是否合法
func ExpValidate(s Exp) bool {
	switch s {
	case EQ, NEQ, GT, GTE, LT, LTE, IN, NOTIN, LIKE, NOTLIKE:
		return true
	default:
		return false
	}
}

// LogicValidate 验证Logic是否合法
func LogicValidate(s Logic) bool {
	switch s {
	case AND, OR:
		return true
	default:
		return false
	}
}

// OrderValidate 验证Order是否合法
func OrderValidate(s Order) bool {
	switch s {
	case ASC, DESC:
		return true
	default:
		return false
	}
}

// ConvertToGormExpression 根据SearchColumn参数转换为符合gorm where clause.Expression
func (p *Req) ConvertToGormExpression(model any) (whereExpressions, orderExpressions []clause.Expression, err error) {
	whereExpressions = make([]clause.Expression, 0)
	orderExpressions = make([]clause.Expression, 0)
	column := fieldToColumn(model)
	if len(p.Query) > 0 {
		for _, v := range p.Query {
			if v.Field == "" {
				return whereExpressions, orderExpressions, fmt.Errorf("field cannot be empty")
			}
			field, ok := column[strings.ToLower(v.Field)]
			if !ok {
				return whereExpressions, orderExpressions, fmt.Errorf("field '%s' is not db column name", v.Field)
			}
			if v.Exp == "" {
				v.Exp = EQ
			}
			if !ExpValidate(v.Exp) {
				return whereExpressions, orderExpressions, fmt.Errorf("unknown exp type '%s'", v.Exp)
			}
			if v.Logic == "" {
				v.Logic = AND
			}
			if !LogicValidate(v.Logic) {
				return whereExpressions, orderExpressions, fmt.Errorf("unknown logic type '%s'", v.Logic)
			}
			var expression clause.Expression
			switch v.Exp {
			case EQ:
				expression = clause.Eq{Column: field, Value: v.Value}
			case NEQ:
				expression = clause.Neq{Column: field, Value: v.Value}
			case GT:
				expression = clause.Gt{Column: field, Value: v.Value}
			case GTE:
				expression = clause.Gte{Column: field, Value: v.Value}
			case LT:
				expression = clause.Lt{Column: field, Value: v.Value}
			case LTE:
				expression = clause.Lte{Column: field, Value: v.Value}
			case IN:
				values, ok := v.Value.([]any)
				if !ok {
					return nil, nil, fmt.Errorf("IN value is not a slice")
				}
				expression = clause.IN{Column: field, Values: values}
			case NOTIN:
				values, ok := v.Value.([]any)
				if !ok {
					return nil, nil, fmt.Errorf("NOTIN value is not a slice")
				}
				expression = clause.Not(clause.IN{Column: field, Values: values})
			case LIKE:
				expression = clause.Like{Column: field, Value: v.Value}
			case NOTLIKE:
				expression = clause.Not(clause.Like{Column: field, Value: v.Value})
			}
			if v.Logic == AND {
				whereExpressions = append(whereExpressions, clause.And(expression))
			} else {
				whereExpressions = append(whereExpressions, clause.Or(expression))
			}
		}
	}
	if len(p.Order) > 0 {
		for _, v := range p.Order {
			if v.Field == "" {
				return whereExpressions, orderExpressions, fmt.Errorf("field cannot be empty")
			}
			field, ok := column[strings.ToLower(v.Field)]
			if !ok {
				return whereExpressions, orderExpressions, fmt.Errorf("field '%s' is not db column name", v.Field)
			}
			if v.Order == "" {
				v.Order = ASC
			}
			if !OrderValidate(v.Order) {
				return whereExpressions, orderExpressions, fmt.Errorf("order is err")
			}
			orderExpressions = append(orderExpressions, clause.OrderBy{
				Columns: []clause.OrderByColumn{
					{
						Column:  clause.Column{Name: field},
						Desc:    v.Order == DESC,
						Reorder: false,
					},
				},
			})
		}
	}
	return whereExpressions, orderExpressions, nil
}

// ConvertToPage 转换为page
func (p *Req) ConvertToPage(total int32) (*Reply, error) {
	resp := &Reply{
		Page:      0,
		PageSize:  0,
		Total:     total,
		PrevPage:  0,
		NextPage:  0,
		TotalPage: 0,
	}
	if p.Page < 0 {
		return resp, fmt.Errorf("page cannot be less than 0")
	}
	if p.PageSize < 0 {
		return resp, fmt.Errorf("pageSize cannot be less than 0")
	}
	if (p.Page != 0 && p.PageSize == 0) || (p.Page == 0 && p.PageSize != 0) {
		return resp, fmt.Errorf("page and pageSize must be a pair")
	}
	if p.Page == 0 && p.PageSize == 0 {
		return resp, nil
	}
	resp.Page = p.Page
	resp.PageSize = p.PageSize
	resp.TotalPage = int32(math.Ceil(float64(total) / float64(p.PageSize)))
	resp.NextPage = p.Page + 1
	if resp.NextPage > resp.TotalPage {
		resp.NextPage = resp.TotalPage
	}
	resp.PrevPage = p.Page - 1
	if resp.PrevPage <= 0 {
		resp.PrevPage = 1
	}
	return resp, nil
}

// fieldToColumn 将model的tag中gorm的tag的Column转换为map[string]string
func fieldToColumn(model any) map[string]string {
	m := make(map[string]string)
	t := reflect.TypeOf(model)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" {
			gormTags := strings.Split(gormTag, ";")
			for _, v := range gormTags {
				if strings.Contains(v, "column") {
					column := strings.Split(v, ":")
					if len(column) == 2 {
						m[strings.ToLower(column[1])] = column[1]
					}
				}
			}
		}
	}
	return m
}
