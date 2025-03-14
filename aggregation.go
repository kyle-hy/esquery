package esquery

// Aggregation 构造聚合查询
func Aggregation(aggName, field, aggType string, params Map) Map {
	aggregation := Map{
		aggName: Map{
			aggType: Map{},
		},
	}

	if params == nil {
		params = Map{}
	}

	aggFields := aggregation[aggName].(Map)[aggType].(Map)

	// 根据不同的聚合类型处理对应的参数
	switch aggType {
	case "terms":
		// 常用于对某个字段进行分桶统计，如按分类、标签等进行分组统计
		aggFields["field"] = field
		if size, ok := params["size"]; ok {
			// 这个参数来调整返回桶的数量,前20个最常见的类别，可以设置 size: 20
			aggFields["size"] = size
		}
		if order, ok := params["order"]; ok {
			// 根据桶的计数（文档数量）或桶内的其他度量（如字段的值）进行排序
			aggFields["order"] = order
		}
		if shardSize, ok := params["shard_size"]; ok {
			// 指定在每个分片上返回多少个桶
			aggFields["shard_size"] = shardSize
		}
	case "range":
		aggFields["field"] = field
		if ranges, ok := params["ranges"]; ok {
			aggFields["ranges"] = ranges
		}
	case "avg", "sum", "max", "min", "cardinality", "stats", "extended_stats", "percentiles", "percentile_ranks":
		aggFields["field"] = field
	case "histogram":
		aggFields["field"] = field
		if interval, ok := params["interval"]; ok {
			aggFields["interval"] = interval
		}
		if min, ok := params["min"]; ok {
			aggFields["min"] = min
		}
		if max, ok := params["max"]; ok {
			aggFields["max"] = max
		}
	case "date_histogram":
		aggFields["field"] = field
		if interval, ok := params["interval"]; ok {
			aggFields["interval"] = interval
		}
		if timeZone, ok := params["time_zone"]; ok {
			aggFields["time_zone"] = timeZone
		}
	case "geo_distance":
		aggFields["field"] = field
		if origin, ok := params["origin"]; ok {
			aggFields["origin"] = origin
		}
		if ranges, ok := params["ranges"]; ok {
			aggFields["ranges"] = ranges
		}
	case "geohash_grid":
		aggFields["field"] = field
		if precision, ok := params["precision"]; ok {
			aggFields["precision"] = precision
		}
	case "filter":
		if query, ok := params["query"]; ok {
			aggFields["query"] = query
		}
		if constantScore, ok := params["constant_score"]; ok {
			aggFields["constant_score"] = constantScore
		}
	case "nested":
		if path, ok := params["path"]; ok {
			aggFields["path"] = path
		}
	case "adjacency_matrix":
		if filters, ok := params["filters"]; ok {
			aggFields["filters"] = filters
		}
	case "top_hits":
		if size, ok := params["size"]; ok {
			aggFields["size"] = size
		}
		if from, ok := params["from"]; ok {
			aggFields["from"] = from
		}
		if sort, ok := params["sort"]; ok {
			aggFields["sort"] = sort
		}
	case "terms_set":
		aggFields["field"] = field
		if minimumShouldMatch, ok := params["minimum_should_match"]; ok {
			// 用于控制查询匹配的文档必须包含多少个查询值
			aggFields["minimum_should_match"] = minimumShouldMatch
		}
	case "bucket_sort":
		if sort, ok := params["sort"]; ok {
			aggFields["sort"] = sort
		}
		if from, ok := params["from"]; ok {
			aggFields["from"] = from
		}
		if size, ok := params["size"]; ok {
			aggFields["size"] = size
		}
	case "scripted_metric":
		if initScript, ok := params["init_script"]; ok {
			aggFields["init_script"] = initScript
		}
		if mapScript, ok := params["map_script"]; ok {
			aggFields["map_script"] = mapScript
		}
		if combineScript, ok := params["combine_script"]; ok {
			aggFields["combine_script"] = combineScript
		}
		if reduceScript, ok := params["reduce_script"]; ok {
			aggFields["reduce_script"] = reduceScript
		}
	case "composite":
		if sources, ok := params["sources"]; ok {
			aggFields["sources"] = sources
		}
		if after, ok := params["after"]; ok {
			aggFields["after"] = after
		}
		if size, ok := params["size"]; ok {
			aggFields["size"] = size
		}
	case "missing":
		aggFields["field"] = field
		if defaultValue, ok := params["default"]; ok {
			aggFields["default"] = defaultValue
		}
	case "sum_of_squares":
		aggFields["field"] = field
		if missing, ok := params["missing"]; ok {
			aggFields["missing"] = missing
		}
		if fields, ok := params["fields"]; ok {
			aggFields["fields"] = fields
		}
	case "bucket_selector":
		if script, ok := params["script"]; ok {
			aggFields["script"] = script
		}
		if bucketsPath, ok := params["buckets_path"]; ok {
			aggFields["buckets_path"] = bucketsPath
		}
	}

	// 添加分页和排序支持
	if sort, ok := params["sort"]; ok {
		aggFields["sort"] = sort
	}
	if size, ok := params["size"]; ok {
		aggFields["size"] = size
	}
	if from, ok := params["from"]; ok {
		aggFields["from"] = from
	}

	return aggregation
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
