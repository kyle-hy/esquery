package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	eq "github.com/kyle-hy/esquery"
)

// TestEsQuery .
func TestEsQuery() {
	// 示例向量
	vector := []float32{0.1, 3.2, 2.1}

	mustQueries := []eq.Map{
		eq.Term("status", "active"),
		eq.Match("title", "Golang开发"),
		eq.Knn("title_vector", vector, nil, eq.WithTopK(5)),
	}

	shouldQueries := []eq.Map{
		eq.Match("description", "快速高效"),
		eq.Match("tags", "编程"),
	}

	filterQueries := []eq.Map{
		eq.Range("price", 100, nil, 200, nil),
		eq.Exists("stock"),
	}

	mustNotQueries := []eq.Map{
		eq.Term("is_deleted", true),
	}

	sort := []eq.Map{
		{"price": map[string]string{"order": "asc"}},
	}

	// 聚合：按类别统计总数
	aggs := eq.TermsAgg("category", eq.WithSize(1))

	// 构造 Elasticsearch 查询
	esQuery := eq.ESQuery{
		Query: eq.Bool(eq.WithMust(mustQueries), eq.WithShould(shouldQueries), eq.WithFilter(filterQueries), eq.WithMustNot(mustNotQueries)),
		Sort:  sort,
		Aggs:  aggs,
	}

	// 输出查询 JSON
	jsonQuery, _ := json.MarshalIndent(esQuery, "", "  ")
	fmt.Println(string(jsonQuery))

}

// ...
func main() {
	// TestEsQuery()

	// 连接数据库查询
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
		},
		APIKey: "",
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// // 构造 Elasticsearch 查询
	esQuery := eq.ESQuery{
		Query: eq.Bool(eq.WithMust([]eq.Map{eq.Match("name", "snow")})),
		Aggs:  eq.TermsAgg("name.keyword", eq.WithSize(2)),
	}
	fmt.Println(esQuery.JSON())

	searchResp, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("books"),
		es.Search.WithBody(esQuery.JSON()),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	fmt.Println(searchResp, err)
}
