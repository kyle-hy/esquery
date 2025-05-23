package esquery

import (
	"fmt"
	"maps"
)

// AggMap 聚合参数的Map
type AggMap Map

// With 组合聚合分析
func (a AggMap) With(m AggMap) AggMap {
	maps.Copy(a, m)
	return a
}

// Nested 嵌套的聚合分析,联合聚合
func (a AggMap) Nested(m AggMap) AggMap {
	a["aggs"] = m
	return a
}

// Aggregation 构造聚合查询（支持 Option 模式）
func Aggregation(field, aggType string, opts ...Option) AggMap {
	aggName := fmt.Sprintf("%s_%s", field, aggType)
	agg := AggMap{
		aggName: AggMap{
			aggType: NewOptMap(opts...), // 使用 Option 组合参数
		},
	}

	// 默认字段参数
	agg[aggName].(AggMap)[aggType].(AggMap)["field"] = field
	return agg
}

// TermsAgg 构造 Terms 聚合查询, 根据字段的不同值进行分组
func TermsAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "terms", opts...)
}

// RangeAgg 构造 Range 聚合查询, 根据指定的范围进行分组
func RangeAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "range", opts...)
}

// AvgAgg 构造 Avg 聚合查询, 计算字段的平均值
func AvgAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "avg", opts...)
}

// SumAgg 构造 Sum 聚合查询, 计算字段的总和
func SumAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "sum", opts...)
}

// MaxAgg 构造 Max 聚合查询, 计算字段的最大值
func MaxAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "max", opts...)
}

// MinAgg 构造 Min 聚合查询, 计算字段的最小值
func MinAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "min", opts...)
}

// CardinalityAgg 构造 Cardinality 聚合查询, 计算字段的去重值数量
func CardinalityAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "cardinality", opts...)
}

// StatsAgg 构造 Stats 聚合查询, 提供字段的统计信息（最小值、最大值、总和、平均值、计数）
func StatsAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "stats", opts...)
}

// ExtendedStatsAgg 构造 Extended Stats 聚合查询, 提供扩展统计信息，包括标准差、方差等
func ExtendedStatsAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "extended_stats", opts...)
}

// PercentilesAgg 构造 Percentiles 聚合查询, 计算字段的百分位数
func PercentilesAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "percentiles", opts...)
}

// PercentileRanksAgg 构造 Percentile Ranks 聚合查询, 计算给定值的百分位数排名
func PercentileRanksAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "percentile_ranks", opts...)
}

// HistogramAgg 构造 Histogram 聚合查询, 按区间对数值字段进行聚合
func HistogramAgg(field string, opts ...Option) AggMap {
	defaultOpts := []Option{WithMinDocCount(0)}
	defaultOpts = append(defaultOpts, opts...)
	return Aggregation(field, "histogram", defaultOpts...)
}

// DateHistogramAgg 构造 Date Histogram 聚合查询, 按日期对字段进行聚合
func DateHistogramAgg(field string, opts ...Option) AggMap {
	defaultOpts := []Option{WithMinDocCount(0)} // 默认函数选项在最前面，若设定则被覆盖
	defaultOpts = append(defaultOpts, opts...)
	return Aggregation(field, "date_histogram", defaultOpts...)
}

// GeoDistanceAgg 构造 Geo Distance 聚合查询, 根据地理位置计算距离分组
func GeoDistanceAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "geo_distance", opts...)
}

// GeohashGridAgg 构造 Geohash Grid 聚合查询, 基于地理位置计算网格聚合
func GeohashGridAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "geohash_grid", opts...)
}

// FilterAgg 构造 Filter 聚合查询, 基于查询条件进行过滤并聚合
func FilterAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "filter", opts...)
}

// NestedAgg 构造 Nested 聚合查询, 用于嵌套查询的聚合
func NestedAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "nested", opts...)
}

// AdjacencyMatrixAgg 构造 Adjacency Matrix 聚合查询, 查找不同字段值之间的关系
func AdjacencyMatrixAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "adjacency_matrix", opts...)
}

// TopHitsAgg 构造 Top Hits 聚合查询, 查询每个桶内的文档，并返回最相关的文档
func TopHitsAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "top_hits", opts...)
}

// TermsSetAgg 构造 Terms Set 聚合查询,根据多个字段的组合进行分组
func TermsSetAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "terms_set", opts...)
}

// BucketSortAgg 构造 Bucket Sort 聚合查询, 对聚合结果桶进行排序
func BucketSortAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "bucket_sort", opts...)
}

// ScriptedMetricAgg 构造 Scripted Metric 聚合查询, 使用自定义脚本执行聚合操作
func ScriptedMetricAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "scripted_metric", opts...)
}

// CompositeAgg 构造 Composite 聚合查询, 允许对多个字段进行聚合，并支持分页
func CompositeAgg(field string, opts ...Option) AggMap {
	return Aggregation(field, "composite", opts...)
}
