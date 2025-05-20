package esquery

import (
	"encoding/json"
	"math"
	"time"
)

// BucketValue 桶聚合的值
type BucketValue struct {
	Value float64 `json:"value"`
}

// DateBucket 日期桶聚合的数据
type DateBucket struct {
	KeyAsString    string             `json:"key_as_string"`
	Key            int64              `json:"key"`
	DocCount       int                `json:"doc_count"`
	DocCountGrowth float64            `json:"doc_count_growth"`      // 文档数量的增长率
	AggsValue      map[string]float64 `json:"aggs_value,omitempty"`  // 动态聚合字段名与值
	AggsGrowth     map[string]float64 `json:"aggs_growth,omitempty"` // 动态聚合字段名与增长率
}

// 解析单个 date_histogram bucket 并提取所有子聚合的数值
func parseDateBucket(data json.RawMessage) (*DateBucket, error) {
	var base struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	}
	if err := json.Unmarshal(data, &base); err != nil {
		return &DateBucket{}, err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return &DateBucket{}, err
	}

	aggsValue := make(map[string]float64)
	for key, val := range raw {
		if key == "key" || key == "key_as_string" || key == "doc_count" {
			continue
		}
		var v BucketValue
		if err := json.Unmarshal(val, &v); err == nil {
			aggsValue[key] = v.Value
		}
	}

	return &DateBucket{
		KeyAsString: base.KeyAsString,
		Key:         base.Key,
		DocCount:    base.DocCount,
		AggsValue:   aggsValue,
	}, nil
}

// getDateHistBuckets 从聚合结果提取结构体
func getDateHistBuckets(aggs map[string]json.RawMessage) map[string][]*DateBucket {
	// 提取值
	buckets := map[string][]*DateBucket{}
	for key, aggData := range aggs {
		aggs := map[string][]json.RawMessage{}
		json.Unmarshal(aggData, &aggs)
		for akey, avalue := range aggs {
			if akey == "buckets" {
				for _, bdata := range avalue {
					db, _ := parseDateBucket(bdata)
					buckets[key] = append(buckets[key], db)
				}
			}
		}
	}
	return buckets
}

// AggsDateHistGrowth 计算聚合增长率(count)
func AggsDateHistGrowth(aggs map[string]json.RawMessage) map[string][]*DateBucket {
	// 提取值
	buckets := getDateHistBuckets(aggs)

	// 计算增长率
	for _, dbs := range buckets {
		preCnt := 0
		for _, value := range dbs {
			value.DocCountGrowth = CntGrowth(preCnt, value.DocCount)
			preCnt = value.DocCount
		}
	}
	return buckets
}

// AggsDateHistStatsGrowth 计算聚合环比增长率
func AggsDateHistStatsGrowth(aggs map[string]json.RawMessage) map[string][]*DateBucket {
	// 提取值
	buckets := getDateHistBuckets(aggs)

	// 计算增长率
	for _, dbs := range buckets {
		var preCnt int
		var preAggs map[string]float64
		for _, value := range dbs {
			value.DocCountGrowth = CntGrowth(preCnt, value.DocCount)
			value.AggsGrowth = CalcGrowth(preAggs, value.AggsValue)
			preCnt = value.DocCount
			preAggs = value.AggsValue
		}
	}
	return buckets
}

// AggsYoYHistStatsGrowth 计算聚合同比增长率
func AggsYoYHistStatsGrowth(aggs map[string]json.RawMessage) map[string][]*DateBucket {
	// 提取值
	buckets := getDateHistBuckets(aggs)

	// 计算增长率
	for _, dbs := range buckets {
		// 转成map结果方便查询和提速
		mapBuckets := map[int64]*DateBucket{}
		for _, db := range dbs {
			mapBuckets[db.Key] = db
		}

		for _, value := range dbs {
			preCnt, preAggs := getPreCntAndAggs(mapBuckets, value)
			value.DocCountGrowth = CntGrowth(preCnt, value.DocCount)
			value.AggsGrowth = CalcGrowth(preAggs, value.AggsValue)
		}
	}
	return buckets
}

// 获取前一个时间的分桶数据
func getPreCntAndAggs(mapBuckets map[int64]*DateBucket, db *DateBucket) (int, map[string]float64) {
	var preCnt int
	var preAggs map[string]float64
	curTime := time.UnixMilli(db.Key)
	preTime := curTime.AddDate(-1, 0, 0)
	pre := mapBuckets[preTime.UnixMilli()]
	if pre != nil {
		preCnt = pre.DocCount
		preAggs = pre.AggsValue
	}
	return preCnt, preAggs
}

// CalcGrowth 计算相邻bucket的增长率
func CalcGrowth(pre, suf map[string]float64) map[string]float64 {
	growth := map[string]float64{}
	// 特殊第一个bucket，增长率都为0
	if len(pre) == 0 {
		for k := range suf {
			growth[k] = 0
		}
		return growth
	}

	// 相邻两个bucket计算增长率
	for k, sufValue := range suf {
		preValue := pre[k]
		if preValue == 0 {
			growth[k] = 0
		} else {
			growth[k] = math.Round(10000*(sufValue-preValue)/preValue) / 10000
		}
	}
	return growth
}

// CntGrowth 计算相邻bucket的数量增长率
func CntGrowth(pre, suf int) float64 {
	if pre == 0 {
		return 0
	}

	return math.Round(10000*float64(suf-pre)/float64(pre)) / 10000
}
