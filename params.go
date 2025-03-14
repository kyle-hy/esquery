package esquery

// CountOrder 按照桶的文档数量进行排序
func CountOrder(orderDirection string) Map {
	return Map{
		"_count": orderDirection, // 按桶的文档数量排序，可以是 "asc" 或 "desc"
	}
}

// TermOrder 按照桶的字段值进行排序
func TermOrder(orderDirection string) Map {
	return Map{
		"_term": orderDirection, // 按桶的字段值排序，可以是 "asc" 或 "desc"
	}
}

// CustomOrder 自定义字段排序
func CustomOrder(field, orderDirection string) Map {
	return Map{
		field: orderDirection, // 按自定义字段排序，可以是 "asc" 或 "desc"
	}
}

// TermsAggParam 用于构造 Terms 聚合查询的参数
// size: 聚合结果返回的桶数量
// shard_size: 每个分片的桶数量
// order: 排序方式，可以是 "_count" 或 "_term"
// min_doc_count: 每个桶的最小文档计数，低于该值的桶会被排除
// include, exclude: 用于限制哪些项被包含或排除，通常用于模式匹配
// missing: 缺失值的聚合方式
func TermsAggParam(size, shardSize int, order Map, minDocCount int, include, exclude string, missing any) Map {
	params := Map{
		"size":          size,        // 返回的桶数量
		"shard_size":    shardSize,   // 每个分片的桶数量
		"order":         order,       // 排序方式，按桶的计数或字段值
		"min_doc_count": minDocCount, // 最小文档计数，低于该值的桶将被排除
	}

	// 可选的 include 和 exclude 参数
	if include != "" {
		params["include"] = include
	}
	if exclude != "" {
		params["exclude"] = exclude
	}
	// 可选的 missing 参数
	if missing != nil {
		params["missing"] = missing
	}

	return params
}
