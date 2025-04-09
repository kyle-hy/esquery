package esquery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
)

// QueryList 查询详情及总数
func QueryList[T any](es *elasticsearch.Client, index string, queryBody any,
) ([]T, int, error) {
	hits, total, _, _, err := QueryWithMeta[T](es, index, queryBody)
	return hits, total, err
}

// RawAgg 必须实现Raw函数的接口
type RawAgg interface {
	Raw(map[string]json.RawMessage) []byte
}

// QueryAgg 查询聚合并将结果解析到指定结构体中
func QueryAgg[T RawAgg](es *elasticsearch.Client, index string, queryBody any) (*T, error) {
	_, _, aggsRaw, _, err := QueryWithMeta[any](es, index, queryBody)
	if err != nil {
		return nil, err
	}

	// 用泛型解析聚合部分
	var result T
	if err := json.Unmarshal(result.Raw(aggsRaw), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// QueryAggRaw 聚合分析查询，返回原始json序列
func QueryAggRaw(es *elasticsearch.Client, index string, queryBody any,
) (map[string]json.RawMessage, error) {
	_, _, aggs, _, err := QueryWithMeta[any](es, index, queryBody)
	return aggs, err
}

// QueryWithMeta 检索及聚合分析结果
func QueryWithMeta[T any](es *elasticsearch.Client, index string, queryBody any,
) ([]T, int, map[string]json.RawMessage, []string, error) {
	queryBytes, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, nil, nil, fmt.Errorf("marshal query failed: %w", err)
	}

	// 执行搜索请求
	res, err := es.Search(
		es.Search.WithContext(context.TODO()),
		es.Search.WithIndex(index),
		es.Search.WithBody(bytes.NewReader(queryBytes)),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, 0, nil, nil, fmt.Errorf("es search failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, 0, nil, nil, fmt.Errorf("es search error: %s", body)
	}

	// 解析响应
	var parsed Result[T]
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return nil, 0, nil, nil, fmt.Errorf("decode response failed: %w", err)
	}

	var results []T
	var ids []string
	for _, hit := range parsed.Hits.Hits {
		results = append(results, hit.Source)
		ids = append(ids, hit.ID)
	}

	return results, parsed.Hits.Total.Value, parsed.Aggregations, ids, nil
}
