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
			Source *T     `json:"_source"`
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

// Raw 提取range对应的json序列
func (t RangeAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_range") {
			return v
		}
	}
	return nil
}

// AvgAggResult 表示avg聚合结果
type AvgAggResult struct {
	Value float64 `json:"value"`
}

// Raw 提取avg对应的json序列
func (r AvgAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_avg") {
			return v
		}
	}
	return nil
}

// SumAggResult 表示sum聚合结果
type SumAggResult struct {
	Value float64 `json:"value"`
}

// Raw 提取sum对应的json序列
func (r SumAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_sum") {
			return v
		}
	}
	return nil
}

// MaxAggResult 表示max聚合结果
type MaxAggResult struct {
	Value float64 `json:"value"`
}

// Raw 提取max对应的json序列
func (r MaxAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_max") {
			return v
		}
	}
	return nil
}

// MinAggResult 表示min聚合结果
type MinAggResult struct {
	Value float64 `json:"value"`
}

// Raw 提取min对应的json序列
func (r MinAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_min") {
			return v
		}
	}
	return nil
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

// Raw 提取date_histogram对应的json序列
func (t DateHistogramAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_date_histogram") {
			return v
		}
	}
	return nil
}

// StatsAggResult 表示统计聚合结果（min/max/avg/sum/count）
type StatsAggResult struct {
	Count int     `json:"count"` // 参与计算的文档数
	Min   float64 `json:"min"`   // 最小值
	Max   float64 `json:"max"`   // 最大值
	Avg   float64 `json:"avg"`   // 平均值
	Sum   float64 `json:"sum"`   // 总和
}

// Raw 提取stats对应的json序列
func (t StatsAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_stats") {
			return v
		}
	}
	return nil
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

// Raw 提取extended_stats对应的json序列
func (t ExtendedStatsAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_extended_stats") {
			return v
		}
	}
	return nil
}

// CardinalityAggResult 表示去重统计结果（唯一值数量）
type CardinalityAggResult struct {
	Value int `json:"value"` // 去重后的值数量
}

// Raw 提取cardinality对应的json序列
func (t CardinalityAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_cardinality") {
			return v
		}
	}
	return nil
}

// ValueCountAggResult 表示字段非空值的文档数
type ValueCountAggResult struct {
	Value int `json:"value"` // 非空字段的文档数量
}

// Raw 提取value对应的json序列
func (t ValueCountAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_value") {
			return v
		}
	}
	return nil
}

// PercentilesAggResult 表示percentiles聚合的结果（百分位数）
type PercentilesAggResult struct {
	Values map[string]float64 `json:"values"`
}

// Raw 提取percentiles对应的json序列
func (r PercentilesAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_percentiles") {
			return v
		}
	}
	return nil
}

// PercentileRanksAggResult 表示 percentile_ranks 聚合结果
type PercentileRanksAggResult struct {
	Values map[string]float64 `json:"values"` // 指定值的百分位排名
}

// Raw 提取 percentile_ranks 聚合的 JSON 数据
func (r PercentileRanksAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_percentile_ranks") {
			return v
		}
	}
	return nil
}

// HistogramAggBucket 表示 histogram 聚合的桶
type HistogramAggBucket struct {
	Key      float64 `json:"key"`       // 区间开始值
	DocCount int     `json:"doc_count"` // 区间内文档数量
}

// HistogramAggResult 表示 histogram 聚合结果
type HistogramAggResult struct {
	Buckets []HistogramAggBucket `json:"buckets"`
}

// Raw 提取 histogram 聚合的 JSON 数据
func (r HistogramAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_histogram") {
			return v
		}
	}
	return nil
}

// GeoDistanceAggBucket 表示 geo_distance 聚合的距离桶
type GeoDistanceAggBucket struct {
	Key      string   `json:"key"`            // 范围标签，如 "0.0-100.0"
	From     *float64 `json:"from,omitempty"` // 起始距离（单位：米）,不一定存在
	To       *float64 `json:"to,omitempty"`   // 结束距离（单位：米）,不一定存在
	DocCount int      `json:"doc_count"`      // 在该范围内的文档数量
}

// FromVal From的值
func (b GeoDistanceAggBucket) FromVal() float64 {
	if b.From != nil {
		return *b.From
	}
	return 0
}

// ToVal To的值
func (b GeoDistanceAggBucket) ToVal() float64 {
	if b.To != nil {
		return *b.To
	}
	return 0
}

// GeoDistanceAggResult 表示 geo_distance 聚合的结果
type GeoDistanceAggResult struct {
	Buckets []GeoDistanceAggBucket `json:"buckets"`
}

// Raw 提取 geo_distance 聚合的 JSON 数据
func (r GeoDistanceAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_geo_distance") {
			return v
		}
	}
	return nil
}

// GeohashGridAggBucket 表示 geohash_grid 聚合的地理网格桶
type GeohashGridAggBucket struct {
	Key      string `json:"key"`       // Geohash 编码
	DocCount int    `json:"doc_count"` // 匹配文档数
}

// GeohashGridAggResult 表示 geohash_grid 聚合的结果
type GeohashGridAggResult struct {
	Buckets []GeohashGridAggBucket `json:"buckets"`
}

// Raw 提取 geohash_grid 聚合的 JSON 数据
func (r GeohashGridAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_geohash_grid") {
			return v
		}
	}
	return nil
}

// FilterAggResult 表示 filter 聚合的结果
type FilterAggResult struct {
	DocCount int `json:"doc_count"` // 满足 filter 条件的文档数量
}

// Raw 提取 filter 聚合的 JSON 数据
func (r FilterAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_filter") {
			return v
		}
	}
	return nil
}

// NestedAggResult 表示 nested 聚合的结果（嵌套文档的 doc_count）
type NestedAggResult struct {
	DocCount int `json:"doc_count"` // 嵌套文档数
}

// Raw 提取 nested 聚合的 JSON 数据
func (r NestedAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_nested") {
			return v
		}
	}
	return nil
}

// AdjacencyMatrixAggBucket 表示 adjacency_matrix 中的每个桶
type AdjacencyMatrixAggBucket struct {
	Key      string `json:"key"`       // 桶的组合键（多个 filter 的交集）
	DocCount int    `json:"doc_count"` // 匹配的文档数
}

// AdjacencyMatrixAggResult 表示 adjacency_matrix 聚合的结果
type AdjacencyMatrixAggResult struct {
	Buckets []AdjacencyMatrixAggBucket `json:"buckets"`
}

// Raw 提取 adjacency_matrix 聚合的 JSON 数据
func (r AdjacencyMatrixAggResult) Raw(agg map[string]json.RawMessage) []byte {
	for k, v := range agg {
		if strings.HasSuffix(k, "_adjacency_matrix") {
			return v
		}
	}
	return nil
}
