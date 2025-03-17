package query

import (
	"strings"
)

var defaultMaxSize = 1000

// SetMaxSize change the default maximum amount pages per page
// 更改默认的每页最大页数
func SetMaxSize(max int) {
	if max < 10 {
		max = 10
	}
	defaultMaxSize = max
}

// Page info
type Page struct {
	page  int // Page number, starting from page 0
	limit int // Number per page
	// Sort fields, default is ID backwards, you can add - sign before the field to indicate reverse order,
	// no - sign to indicate ascending order, multiple fields separated by a comma.
	sort string
}

// Page get page value
func (p *Page) Page() int {
	return p.page
}

// Limit number per page
func (p *Page) Limit() int {
	return p.limit
}

// Size number per page
// Deprecated: use Limit instead
func (p *Page) Size() int {
	return p.limit
}

// Sort get sort field
func (p *Page) Sort() string {
	return p.sort
}

// Offset get offset value
func (p *Page) Offset() int {
	return p.page * p.limit
}

// DefaultPage default page, number 20 per page, sorted by id backwards
// 默认页面，每页 20 个，按 ID 倒序排序
func DefaultPage(page int) *Page {
	if page < 0 {
		page = 0
	}
	return &Page{
		page:  page,
		limit: 20,
		sort:  "id DESC",
	}
}

// NewPage 创建一个新的分页对象
// 参数:
// - page: 页码，从0开始
// - limit: 每页记录数
// - columnNames: 排序字段，多个字段用逗号分隔，字段前加'-'表示降序
func NewPage(page int, limit int, columnNames string) *Page {
	if page < 0 {
		page = 0
	}
	if limit > defaultMaxSize || limit < 1 { // 请求过大的数据量导致性能问题
		limit = defaultMaxSize
	}

	return &Page{
		page:  page,
		limit: limit,
		sort:  getSort(columnNames),
	}
}

// convert to mysql sort, each column name preceded by a '-' sign,
// indicating descending order, otherwise ascending order, example:
// 转换为 MySQL 排序，每个列名前面都有一个 '-' 符号，表示降序，否则为升序，示例：
//
//	 getSort 根据 columnNames 生成排序子句
//		columnNames="name" means sort by name in ascending order, 表示按名称升序排序，
//		columnNames="-name" means sort by name descending,
//		columnNames="name,age" means sort by name in ascending order, otherwise sort by age in ascending order,
//		columnNames="-name,-age" means sort by name descending before sorting by age descending.
func getSort(columnNames string) string {
	columnNames = strings.Replace(columnNames, " ", "", -1)
	if columnNames == "" {
		return "id DESC"
	}

	names := strings.Split(columnNames, ",")
	strs := make([]string, 0, len(names))
	for _, name := range names {
		if name[0] == '-' && len(name) > 1 {
			strs = append(strs, name[1:]+" DESC")
		} else {
			strs = append(strs, name+" ASC")
		}
	}

	return strings.Join(strs, ", ")
}
