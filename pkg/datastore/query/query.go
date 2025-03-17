package query

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// Eq equal
	Eq = "eq"
	// Neq not equal
	Neq = "neq"
	// Gt greater than
	Gt = "gt"
	// Gte greater than or equal
	Gte = "gte"
	// Lt less than
	Lt = "lt"
	// Lte less than or equal
	Lte = "lte"
	// Like fuzzy lookup
	Like = "like"
	// In include
	In = "in"
	// NotIN not include
	NotIN = "notin"
	// IsNull is null
	IsNull = "isnull"
	// IsNotNull is not null
	IsNotNull = "is notnull"

	// AND logic and
	AND string = "and"
	// OR logic or
	OR string = "or"
)

var expMap = map[string]string{
	Eq:        " = ",
	Neq:       " <> ",
	Gt:        " > ",
	Gte:       " >= ",
	Lt:        " < ",
	Lte:       " <= ",
	Like:      " LIKE ",
	In:        " IN ",
	NotIN:     " NOT IN ",
	IsNull:    " IS NULL ",
	IsNotNull: " IS NOT NULL ",

	"=":           " = ",
	"!=":          " <> ",
	">":           " > ",
	">=":          " >= ",
	"<":           " < ",
	"<=":          " <= ",
	"not in":      " NOT IN ",
	"is null":     " IS NULL ",
	"is not null": " IS NOT NULL ",
}

var logicMap = map[string]string{
	AND: " AND ",
	OR:  " OR ",

	"&":   " AND ",
	"&&":  " AND ",
	"|":   " OR ",
	"||":  " OR ",
	"AND": " AND ",
	"OR":  " OR ",

	"and:(": " AND ",
	"and:)": " AND ",
	"or:(":  " OR ",
	"or:)":  " OR ",
}

// Params query parameters 查询参数
type Params struct {
	Page       int         `json:"page" form:"page" binding:"gte=0"`
	Limit      int         `json:"limit" form:"limit" binding:"gte=1"`
	Sort       string      `json:"sort,omitempty" form:"sort" binding:""`
	RowColumns interface{} `json:"row_columns,omitempty" form:"row_columns" binding:""`

	Columns []Column `json:"columns,omitempty" form:"columns"` // Not required
}

// Column query info 查询信息
type Column struct {
	Name  string      `json:"name" form:"name"`   // column name
	Exp   string      `json:"exp" form:"exp"`     // expressions, default value is "=", support =, !=, >, >=, <, <=, like, in, notin, isnull, is notnull
	Value interface{} `json:"value" form:"value"` // column value
	Logic string      `json:"logic" form:"logic"` // Logical type, defaults to and when the value is null, with &(and), ||(or)
}

func (c *Column) checkValid() error {
	if c.Name == "" {
		return fmt.Errorf("field 'name' cannot be empty")
	}
	if c.Value == nil {
		v := expMap[strings.ToLower(c.Exp)]
		if v == " IS NULL " || v == " IS NOT NULL " {
			return nil
		}
		return fmt.Errorf("field 'value' cannot be nil")
	}
	return nil
}

// Converting ExpType to sql expressions and LogicType to sql using characters
// 使用字符将 ExpType 转换为 sql 表达式, 将 LogicType 转换为 sql.
func (c *Column) convert() (string, error) {
	symbol := "?"
	if c.Exp == "" {
		c.Exp = Eq
	}
	if v, ok := expMap[strings.ToLower(c.Exp)]; ok { // nolint
		c.Exp = v
		switch c.Exp {
		case " LIKE ":
			val, ok1 := c.Value.(string)

			if !ok1 {
				return symbol, fmt.Errorf("invalid value type '%s'", c.Value)
			}
			l := len(val)
			if l > 2 {
				val2 := val[1 : l-1]
				val2 = strings.ReplaceAll(val2, "%", "\\%")
				val2 = strings.ReplaceAll(val2, "_", "\\_")
				val = string(val[0]) + val2 + string(val[l-1])
			}
			if strings.HasPrefix(val, "%") ||
				strings.HasPrefix(val, "_") ||
				strings.HasSuffix(val, "%") ||
				strings.HasSuffix(val, "_") {
				c.Value = val
			} else {
				c.Value = "%" + val + "%"
			}
		case " IN ", " NOT IN ":
			val, ok1 := c.Value.(string)
			if !ok1 {
				return symbol, fmt.Errorf("invalid value type '%s'", c.Value)
			}
			var iVal []interface{}
			ss := strings.Split(val, ",")
			for _, s := range ss {
				iVal = append(iVal, s)
			}
			c.Value = iVal
			symbol = "(?)"
		case " IS NULL ", " IS NOT NULL ":
			c.Value = nil
			symbol = ""
		}
	} else {
		return symbol, fmt.Errorf("unsported exp type '%s'", c.Exp)
	}

	if c.Logic == "" {
		c.Logic = AND
	} else {
		logic := strings.ToLower(c.Logic)
		if _, ok := logicMap[logic]; ok { // nolint
			c.Logic = logic
		} else {
			return symbol, fmt.Errorf("unsported logic type '%s'", c.Logic)
		}
	}

	return symbol, nil
}

// ConvertToPage converted to page 转化为sql参数
func (p *Params) ConvertToPage() (order string, limit int, offset int) {
	page := NewPage(p.Page, p.Limit, p.Sort)
	order = page.sort
	limit = page.limit
	offset = page.page * page.limit
	return
}

// ConvertToGormConditions conversion to gorm-compliant parameters based on the Columns parameter
// ignore the logical type of the last column, whether it is a one-column or multi-column query
// 根据 Columns 参数转换为符合 gorm 的参数会忽略最后一列的逻辑类型, 无论是单列查询还是多列查询.
func (p *Params) ConvertToGormConditions() (string, []interface{}, error) {
	err := p.ConvertToStructColumn()
	if err != nil {
		return "", nil, err
	}
	str := ""
	var args []interface{}
	l := len(p.Columns)
	if l == 0 {
		return "", nil, nil
	}

	isUseIN := true
	if l == 1 {
		isUseIN = false
	}
	field := p.Columns[0].Name

	for i, column := range p.Columns {
		if err := column.checkValid(); err != nil {
			return "", nil, err
		}

		symbol, err := column.convert()
		if err != nil {
			return "", nil, err
		}

		if i == l-1 { // ignore the logical type of the last column
			switch column.Logic {
			case "or:)", "and:)":
				str += column.Name + column.Exp + symbol + " ) "
			default:
				str += column.Name + column.Exp + symbol
			}
		} else {
			switch column.Logic {
			case "or:(", "and:(":
				str += " ( " + column.Name + column.Exp + symbol + logicMap[column.Logic]
			case "or:)", "and:)":
				str += column.Name + column.Exp + symbol + " ) " + logicMap[column.Logic]
			default:
				str += column.Name + column.Exp + symbol + logicMap[column.Logic]
			}
		}
		if column.Value != nil {
			args = append(args, column.Value)
		}
		// when multiple columns are the same, determine whether [OMIT] IN.
		if isUseIN {
			if field != column.Name {
				isUseIN = false
				continue
			}
			if column.Exp != expMap[Eq] {
				isUseIN = false
			}
		}
	}

	if isUseIN {
		str = field + " IN (?)"
		args = []interface{}{args}
	}

	return str, args, nil
}

func (p *Params) ConvertToStructColumn() error {
	// 检查 RowColumns 是否为结构体
	rowColumnsValue := reflect.ValueOf(p.RowColumns)
	if rowColumnsValue.Kind() != reflect.Struct {
		return fmt.Errorf("RowColumns must be a map struct")
	}

	// 获取结构体的类型信息
	rowStructType := reflect.ValueOf(p.RowColumns)

	tagToValue := make(map[string]interface{})
	for i := 0; i < rowStructType.NumField(); i++ {
		value := rowStructType.Field(i)
		tags := rowStructType.Type().Field(i).Tag
		tagToValue[tags.Get("json")] = value.Interface()
	}

	// 遍历 Columns 数组，根据 Name 字段匹配 RowColumns 的字段
	for i := range p.Columns {
		// 获取 Column.Name
		column := &p.Columns[i] // 获取原始元素的指针
		name := column.Name
		// 查找结构体中 JSON 标签匹配的字段
		if value, ok := tagToValue[name]; ok {
			column.Value = value
		}
	}
	return nil
}

// Conditions query conditions
type Conditions struct {
	Columns []Column `json:"columns" form:"columns" binding:"min=1"` // columns info
}

// ConvertToGorm conversion to gorm-compliant parameters based on the Columns parameter
// ignore the logical type of the last column, whether it is a one-column or multi-column query
// 根据 Columns 参数转换为 gorm 兼容参数会忽略最后一列的逻辑类型, 无论是单列查询还是多列查询.
func (c *Conditions) ConvertToGorm() (string, []interface{}, error) {
	p := &Params{Columns: c.Columns}
	return p.ConvertToGormConditions()
}

// CheckValid check valid
func (c *Conditions) CheckValid() error {
	if len(c.Columns) == 0 {
		return fmt.Errorf("field 'columns' cannot be empty")
	}
	return nil
}
