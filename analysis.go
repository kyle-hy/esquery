package esquery

import "encoding/json"

// 聚合分析类型
const (
	DateHistGrowth      = "dateHistGrowth"      // 按日期分桶统计数量后的增长率分析
	DateHistStatsGrowth = "dateHistStatsGrowth" // 按日期分桶后数值字段统计的增长率分析
	YoYHistStatsGrowth = "yoyHistStatsGrowth" // 按日期分桶后数值字段统计的增长率分析
)

// AggsAnalysis 对聚合结果进行增长率等分析
func AggsAnalysis(aggs map[string]json.RawMessage, atype string) any {
	switch atype {
	case DateHistGrowth:
		return AggsDateHistGrowth(aggs)
	case DateHistStatsGrowth:
		return AggsDateHistStatsGrowth(aggs)
	case YoYHistStatsGrowth:
		return AggsYoYHistStatsGrowth(aggs)
	}
	return aggs
}
