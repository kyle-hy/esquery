package main

import (
	"encoding/json"
	"fmt"

	ep "github.com/kyle-hy/esquery"
)

func main() {

	// 示例向量
	vector := []float32{0.1, 3.2, 2.1}

	mustQueries := []ep.Map{
		ep.Term("status", "active"),
		ep.Match("title", "Golang开发"),
		ep.Knn("title_vector", vector, nil, ep.WithTopK(5)),
	}

	shouldQueries := []ep.Map{
		ep.Match("description", "快速高效"),
		ep.Match("tags", "编程"),
	}

	filterQueries := []ep.Map{
		ep.Range("price", 100, nil, 200, nil),
		ep.Exists("stock"),
	}

	mustNotQueries := []ep.Map{
		ep.Term("is_deleted", true),
	}

	sort := []ep.Map{
		{"price": map[string]string{"order": "asc"}},
	}

	// 聚合：按类别统计总数
	// aggs := ep.Aggregation("category", "terms", ep.WithSize(1))
	aggs := ep.TermsAgg("category", ep.WithSize(1))

	// 构造 Elasticsearch 查询
	esQuery := ep.ESQuery{
		Index: "products",
		Query: ep.Bool(mustQueries, shouldQueries, filterQueries, mustNotQueries),
		Sort:  sort,
		From:  0,
		Size:  10,
		Aggs:  aggs,
	}

	// 输出查询 JSON
	jsonQuery, _ := json.MarshalIndent(esQuery, "", "  ")
	fmt.Println(string(jsonQuery))
}
