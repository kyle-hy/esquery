package esquery

import (
	"encoding/json"
	"strings"
)

// Result es的查询结果解析
type Result[T any] struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string `json:"_id"`
			Source T      `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Aggregations map[string]json.RawMessage `json:"aggregations"` // json.RawMessage 使用各AggResult解析
}

// TermsAggBucket 表示 terms 聚合中的一个桶（Bucket）
// 每个 bucket 表示一个唯一的 term 及其文档数量
type TermsAggBucket struct {
	Key      string `json:"key"`       // 聚合键值（term）
	DocCount int    `json:"doc_count"` // 匹配该 term 的文档数
}

// TermsAggResult 表示 terms 聚合的整体结果
type TermsAggResult struct {
	Buckets []TermsAggBucket `json:"buckets"` // 所有 bucket 列表
}

// Raw 提取terms对应的json序列
func (t TermsAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_terms") {
			return v
		}
	}
	return nil
}

// RangeAggBucket 表示 range 聚合中的一个范围桶
type RangeAggBucket struct {
	Key      string   `json:"key"`            // 自定义 key（例如 "0-100"）
	From     *float64 `json:"from,omitempty"` // 范围起始（可为 nil）
	To       *float64 `json:"to,omitempty"`   // 范围结束（可为 nil）
	DocCount int      `json:"doc_count"`      // 属于该范围的文档数量
}

// RangeAggResult 表示 range 聚合的结果
type RangeAggResult struct {
	Buckets []RangeAggBucket `json:"buckets"` // 所有范围 bucket
}

// DateHistogramBucket 表示某个时间区间的统计
type DateHistogramBucket struct {
	KeyAsString string `json:"key_as_string"` // 可读的时间字符串（如 "2024-01-01T00:00:00Z"）
	Key         int64  `json:"key"`           // 毫秒时间戳
	DocCount    int    `json:"doc_count"`     // 区间内文档数
}

// DateHistogramAggResult 表示 date_histogram 聚合结果
type DateHistogramAggResult struct {
	Buckets []DateHistogramBucket `json:"buckets"` // 时间区间桶
}

// StatsAggResult 表示统计聚合结果（min/max/avg/sum/count）
type StatsAggResult struct {
	Count int     `json:"count"` // 参与计算的文档数
	Min   float64 `json:"min"`   // 最小值
	Max   float64 `json:"max"`   // 最大值
	Avg   float64 `json:"avg"`   // 平均值
	Sum   float64 `json:"sum"`   // 总和
}

// ExtendedStatsAggResult 提供更详细的统计信息
type ExtendedStatsAggResult struct {
	Count        int     `json:"count"`          // 文档数
	Min          float64 `json:"min"`            // 最小值
	Max          float64 `json:"max"`            // 最大值
	Avg          float64 `json:"avg"`            // 平均值
	Sum          float64 `json:"sum"`            // 总和
	SumOfSquares float64 `json:"sum_of_squares"` // 平方和
	Variance     float64 `json:"variance"`       // 方差
	StdDeviation float64 `json:"std_deviation"`  // 标准差
}

// CardinalityAggResult 表示去重统计结果（唯一值数量）
type CardinalityAggResult struct {
	Value int `json:"value"` // 去重后的值数量
}

// ValueCountAggResult 表示字段非空值的文档数
type ValueCountAggResult struct {
	Value int `json:"value"` // 非空字段的文档数量
}
