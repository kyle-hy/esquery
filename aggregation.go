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
