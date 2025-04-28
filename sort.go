package esquery

// 排序方式
const (
	OrderDesc = "desc"
	OrderAsc  = "asc"
)

// SortMap 指定指定的排序Map列表
type SortMap []Map

// With 组合聚合分析
func (s SortMap) With(field, order string) SortMap {
	s = append(s, Sort(field, order)...)
	return s
}

// Sort 查询结果排序
func Sort(field, order string) SortMap {
	s := Map{field: Map{"order": order}}
	return SortMap{s}
}
