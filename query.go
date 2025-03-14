package esquery

// Map 用于表示 Elasticsearch 查询的键值对
type Map = map[string]any

// ESQuery 定义主查询结构
type ESQuery struct {
	Index string `json:"index"`          // 索引名
	Query Map    `json:"query"`          // 查询条件
	Sort  []Map  `json:"sort,omitempty"` // 排序条件
	From  int    `json:"from,omitempty"` // 分页起始位置
	Size  int    `json:"size,omitempty"` // 每页返回条数
	Aggs  Map    `json:"aggs,omitempty"` // 聚合条件
}

// TermQuery 构造 Term 查询, 进行精确匹配
// @param field 查询字段
// @param value 查询值, 不进行分词
// @param opts option不定参数
func TermQuery(field string, value any, opts ...Option) Map {
	paramMap := newOptMap(opts)
	paramMap[field] = value
	return Map{"term": paramMap}
}

// TermsQuery 构造 Term 查询, 进行精确匹配
// @param field 查询字段
// @param values 查询值不定参数列表, 不进行分词
// @param opts option不定参数
func TermsQuery(field string, values []any, opts ...Option) Map {
	paramMap := newOptMap(opts)
	paramMap[field] = values
	return Map{"terms": paramMap}
}

// MatchQuery 构造 Match 查询, 全文检索(模糊匹配、分词搜索、相关度评分)
// @param field 查询字段
// @param value 查询值, 进行分词
func MatchQuery(field string, value any, opts ...Option) Map {
	paramMap := newOptMap(opts)
	paramMap[field] = value
	return Map{"match": paramMap}
}

// MultiMatchQuery 构造MultiMatch查询, match查询的多字段版本，专为多字段搜索设计
func MultiMatchQuery(query string, fields []string, opts ...Option) Map {
	paramMap := newOptMap(opts)
	paramMap["query"] = query
	paramMap["fields"] = fields
	return Map{"multi_match": paramMap}
}

// RangeQuery 构造 Range 查询
func RangeQuery(field string, gte, lte any, opts ...Option) Map {
	paramMap := newOptMap(opts)
	paramMap["gte"] = gte
	paramMap["lte"] = lte
	return Map{
		"range": Map{
			field: paramMap,
		},
	}
}

// NestedQuery 构造 Nested 查询
func NestedQuery(path string, query Map) Map {
	return Map{
		"nested": Map{
			"path":  path,
			"query": query,
		},
	}
}

// ScriptScoreQuery 构造 Script Score 查询
func ScriptScoreQuery(script string) Map {
	return Map{
		"script_score": Map{
			"script": Map{
				"source": script,
			},
		},
	}
}

// WildcardQuery 构造 Wildcard 查询
func WildcardQuery(field string, value string) Map {
	return Map{
		"wildcard": Map{
			field: value,
		},
	}
}

// ExistsQuery 构造 Exists 查询
func ExistsQuery(field string) Map {
	return Map{
		"exists": Map{
			"field": field,
		},
	}
}

// GeoDistanceQuery 构造 GeoDistance 查询
func GeoDistanceQuery(field string, lat, lon float64, distance string) Map {
	return Map{
		"geo_distance": Map{
			"distance": distance,
			field: map[string]float64{
				"lat": lat,
				"lon": lon,
			},
		},
	}
}

// KnnQuery 构造 KNN 查询
// @param field 查询字段
// @param vector 查询向量
// @param topK 返回最相似的结果数
// @param numCandidates 候选数量
// @param boost KNN 查询的 boost 值
func KnnQuery(field string, vector []float32, topK int, numCandidates int, boost float64) Map {
	return Map{
		"knn": Map{
			"field":          field,
			"query_vector":   vector,
			"k":              topK,
			"num_candidates": numCandidates,
			"boost":          boost, // 添加 boost
		},
	}
}

// BoolQuery 构造 Bool 查询（支持 must、should、filter、must_not、minimum_should_match、boost）
// @param must                  必须满足的查询
// @param should                可选查询（至少满足`minimumShouldMatch`个）
// @param filter                过滤器（不影响评分）
// @param mustNot               必须不满足的查询
// @param minimumShouldMatch    `should` 查询中最少满足的条件数
// @param boost                 查询的权重评分
func BoolQuery(must, should, filter, mustNot []Map, minimumShouldMatch int, boost float64) Map {
	boolQuery := Map{}

	if len(must) > 0 {
		boolQuery["must"] = must
	}
	if len(should) > 0 {
		boolQuery["should"] = should
	}
	if len(filter) > 0 {
		boolQuery["filter"] = filter
	}
	if len(mustNot) > 0 {
		boolQuery["must_not"] = mustNot
	}
	if minimumShouldMatch > 0 {
		// 用于控制查询匹配的文档必须包含多少个查询值
		boolQuery["minimum_should_match"] = minimumShouldMatch
	}
	if boost != 0 {
		boolQuery["boost"] = boost
	}

	return Map{
		"bool": boolQuery,
	}
}

// TermsAgg 构造 Terms 聚合查询, 根据字段的不同值进行分组
func TermsAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "terms", params)
}

// RangeAgg 构造 Range 聚合查询, 根据指定的范围进行分组
func RangeAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "range", params)
}

// AvgAgg 构造 Avg 聚合查询, 计算字段的平均值
func AvgAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "avg", params)
}

// SumAgg 构造 Sum 聚合查询, 计算字段的总和
func SumAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "sum", params)
}

// MaxAgg 构造 Max 聚合查询, 计算字段的最大值
func MaxAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "max", params)
}

// MinAgg 构造 Min 聚合查询, 计算字段的最小值
func MinAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "min", params)
}

// CardinalityAgg 构造 Cardinality 聚合查询, 计算字段的去重值数量
func CardinalityAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "cardinality", params)
}

// StatsAgg 构造 Stats 聚合查询, 提供字段的统计信息（最小值、最大值、总和、平均值、计数）
func StatsAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "stats", params)
}

// ExtendedStatsAgg 构造 Extended Stats 聚合查询, 提供扩展统计信息，包括标准差、方差等
func ExtendedStatsAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "extended_stats", params)
}

// PercentilesAgg 构造 Percentiles 聚合查询, 计算字段的百分位数
func PercentilesAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "percentiles", params)
}

// PercentileRanksAgg 构造 Percentile Ranks 聚合查询, 计算给定值的百分位数排名
func PercentileRanksAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "percentile_ranks", params)
}

// HistogramAgg 构造 Histogram 聚合查询, 按区间对数值字段进行聚合
func HistogramAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "histogram", params)
}

// DateHistogramAgg 构造 Date Histogram 聚合查询, 按日期对字段进行聚合
func DateHistogramAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "date_histogram", params)
}

// GeoDistanceAgg 构造 Geo Distance 聚合查询, 根据地理位置计算距离分组
func GeoDistanceAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "geo_distance", params)
}

// GeohashGridAgg 构造 Geohash Grid 聚合查询, 基于地理位置计算网格聚合
func GeohashGridAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "geohash_grid", params)
}

// FilterAgg 构造 Filter 聚合查询, 基于查询条件进行过滤并聚合
func FilterAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "filter", params)
}

// NestedAgg 构造 Nested 聚合查询, 用于嵌套查询的聚合
func NestedAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "nested", params)
}

// AdjacencyMatrixAgg 构造 Adjacency Matrix 聚合查询, 查找不同字段值之间的关系
func AdjacencyMatrixAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "adjacency_matrix", params)
}

// TopHitsAgg 构造 Top Hits 聚合查询, 查询每个桶内的文档，并返回最相关的文档
func TopHitsAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "top_hits", params)
}

// TermsSetAgg 构造 Terms Set 聚合查询,根据多个字段的组合进行分组
func TermsSetAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "terms_set", params)
}

// BucketSortAgg 构造 Bucket Sort 聚合查询, 对聚合结果桶进行排序
func BucketSortAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "bucket_sort", params)
}

// ScriptedMetricAgg 构造 Scripted Metric 聚合查询, 使用自定义脚本执行聚合操作
func ScriptedMetricAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "scripted_metric", params)
}

// CompositeAgg 构造 Composite 聚合查询, 允许对多个字段进行聚合，并支持分页
func CompositeAgg(aggName, field string, params Map) Map {
	return Aggregation(aggName, field, "composite", params)
}
