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
		ep.TermQuery("status", "active"),
		ep.MatchQuery("title", "Golang开发"),
		ep.KnnQuery("title_vector", vector, 5, 100, 2.0),
	}

	shouldQueries := []ep.Map{
		ep.MatchQuery("description", "快速高效"),
		ep.MatchQuery("tags", "编程"),
	}

	filterQueries := []ep.Map{
		ep.RangeQuery("price", 100, nil, 200, nil),
		ep.ExistsQuery("stock"),
	}

	mustNotQueries := []ep.Map{
		ep.TermQuery("is_deleted", true),
	}

	sort := []ep.Map{
		{"price": map[string]string{"order": "asc"}},
	}

	// 聚合：按类别统计总数
	aggs := ep.Map{
		"category_count": ep.Aggregation("category_count", "category", "terms", nil),
	}

	// 构造 Elasticsearch 查询
	esQuery := ep.ESQuery{
		Index: "products",
		Query: ep.BoolQuery(mustQueries, shouldQueries, filterQueries, mustNotQueries, 1, 1.5),
		Sort:  sort,
		From:  0,
		Size:  10,
		Aggs:  aggs,
	}

	// 输出查询 JSON
	jsonQuery, _ := json.MarshalIndent(esQuery, "", "  ")
	fmt.Println(string(jsonQuery))
}
