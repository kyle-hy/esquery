package esquery

import (
	"bytes"
	"encoding/json"
)

// Map 用于表示 Elasticsearch 查询的键值对
type Map = map[string]any

// QueryInfo 查询信息，用于重入如分页查询等
type QueryInfo struct {
	Index string // 查询的索引名
	Query string // 查询语句DSL
}

// ESQuery 定义主查询结构
type ESQuery struct {
	Query Map   `json:"query"`          // 查询条件
	Sort  []Map `json:"sort,omitempty"` // 排序条件
	Aggs  Map   `json:"aggs,omitempty"` // 聚合条件
}

// JSON json序列化
func (eq *ESQuery) JSON() *bytes.Buffer {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(eq)
	return &buf
}

// Bool 构造Bool查询（支持 must、should、filter、must_not等）
func Bool(opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	return Map{"bool": paramMap}
}

// Term 构造 Term 查询, 进行精确匹配
// @param field 查询字段
// @param value 查询值, 不进行分词
// @param opts option不定参数
func Term(field string, value any, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap[field] = value
	return Map{"term": paramMap}
}

// Terms 构造 Term 查询, 进行精确匹配
// @param field 查询字段
// @param values 查询值不定参数列表, 不进行分词
// @param opts option不定参数
func Terms(field string, values []any, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap[field] = values
	return Map{"terms": paramMap}
}

// Match 构造 Match 查询, 全文检索(模糊匹配、分词搜索、相关度评分)
// @param field 查询字段
// @param value 查询值, 进行分词
// @param opts option不定参数
func Match(field string, value any, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap[field] = value
	return Map{"match": paramMap}
}

// MultiMatch 构造MultiMatch查询, match查询的多字段版本，专为多字段搜索设计
// @param query 查询值
// @param fields 查询字段列表
// @param opts option不定参数
func MultiMatch(query string, fields []string, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap["query"] = query
	paramMap["fields"] = fields
	return Map{"multi_match": paramMap}
}

// Range 构造 Range 查询
// @param field 查询字段
// @param gte (>=)大于等于指定值
// @param gt (>)大于指定值
// @param lt (<)小于指定值
// @param lte (<=)小于等于指定值
// @param opts option不定参数
func Range(field string, gte, gt, lt, lte any, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	if gte != nil {
		paramMap["gte"] = gte
	}
	if gt != nil {
		paramMap["gt"] = gt
	}
	if lt != nil {
		paramMap["lt"] = lt
	}
	if lte != nil {
		paramMap["lte"] = lte
	}
	return Map{
		"range": Map{
			field: paramMap,
		},
	}
}

// Nested 构造Nested查询, 用于在嵌套字段（nested type）文档中执行独立的查询
// @param path 索引路径(名称)
// @param query 查询语句
// @param opts option不定参数
func Nested(path string, query Map, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap["path"] = path
	paramMap["query"] = query
	return Map{"nested": paramMap}
}

// ScriptScore 构造Script Score查询, 通过自定义脚本计算每个文档的得分
// @param query 查询语句
// @param script 计算score脚本
// @param opts option不定参数, 为script指定参数
func ScriptScore(query Map, script string, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	return Map{
		"script_score": Map{
			"query": query,
			"script": Map{
				"source": script,
				"params": paramMap,
			},
		},
	}
}

// Wildcard 构造Wildcard查询, 基于通配符的字符串匹配
// @param field 查询字段
// @param value 通配符表达式
// @param opts option不定参数, 为script指定参数
func Wildcard(field string, value string, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap["value"] = value
	return Map{
		"wildcard": Map{
			field: paramMap,
		},
	}
}

// Exists 构造Exists查询,判断某个字段是否存在
// @param field 判断的字段
func Exists(field string) Map {
	return Map{
		"exists": Map{
			"field": field,
		},
	}
}

// GeoDistance 构造 GeoDistance 查询
// @param field 查询字段
// @param lat 纬度
// @param lon 经度
// @param distance 距离阈值, 带单位km、mi、m、yd、ft
// @param opts option不定参数, 为script指定参数
func GeoDistance(field string, lat, lon float64, distance string, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap["distance"] = distance
	paramMap[field] = map[string]float64{"lat": lat, "lon": lon}
	return Map{"geo_distance": paramMap}
}

// Knn 构造 KNN 查询
// @param field 查询字段
// @param vector 查询向量
// @param filter 过滤条件
// @param opts option不定参数
func Knn(field string, vector []float32, opts ...Option) Map {
	paramMap := NewOptMap(opts...)
	paramMap["field"] = field
	paramMap["query_vector"] = vector
	return Map{"knn": paramMap}
}
