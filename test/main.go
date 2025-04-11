package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	eq "github.com/kyle-hy/esquery"
)

// JPrint json序列化后终端打印
func JPrint(data any) {
	jdata, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println("json:", string(jdata))
}

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

// StringInt create a type alias for type int
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
// / this helps us check the type of our value before unmarshalling it
func (st *StringInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into a int we convert it
	///otherwise we return an error
	var item any
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.Atoi(v)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StringInt(i)

	}
	return nil
}

// 日期类型
var (
	ErrInvalidDateFormat = errors.New("invalid date format")
)

// SmartDate 支持多种格式和时间戳自动解析的时间类型（适配 date 和 date_nanos）
type SmartDate struct {
	time.Time
}

// 支持的字符串格式（按优先顺序）
var supportedFormats = []string{
	time.RFC3339Nano,      // 2006-01-02T15:04:05.999999999Z07:00
	time.RFC3339,          // 2006-01-02T15:04:05Z07:00
	"2006-01-02",          // 1992-06-01
	"2006/01/02",          // 1992/06/01
	"2006-01-02 15:04:05", // 1992-06-01 12:30:45
}

// UnmarshalJSON 自动支持多种时间格式和时间戳（纳秒、毫秒、秒）
func (sd *SmartDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		return nil
	}

	// 尝试解析为整数时间戳（可能是秒/ms/ns）
	if ts, err := strconv.ParseInt(s, 10, 64); err == nil {
		switch len(s) {
		case 10: // 秒
			sd.Time = time.Unix(ts, 0)
		case 13: // 毫秒
			sd.Time = time.UnixMilli(ts)
		case 16: // 微秒
			sd.Time = time.Unix(0, ts*1e3)
		case 19: // 纳秒
			sd.Time = time.Unix(0, ts)
		default:
			return errors.New("unsupported numeric timestamp length")
		}
		return nil
	}

	// 依次尝试已知格式
	for _, layout := range supportedFormats {
		if t, err := time.Parse(layout, s); err == nil {
			sd.Time = t
			return nil
		}
	}

	return errors.New("unsupported time format: " + s)
}

// MarshalJSON 序列化为 RFC3339 格式字符串
func (sd SmartDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(sd.Time.Format(time.RFC3339))
}

// String 返回可读格式
func (sd SmartDate) String() string {
	return sd.Time.Format("2006-01-02")
}

// Books .
type Books struct {
	Name        string    `json:"name"`         // 书名
	Author      string    `json:"author"`       // 作者
	ReleaseDate SmartDate `json:"release_date"` // 出版日期
	PageCount   StringInt `json:"page_count"`   // 页数
}

// ...
func main() {
	// TestEsQuery()

	// 连接数据库查询
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
			// "https://192.168.33.1:9200/",
		},
		APIKey: "",
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// // 构造 Elasticsearch 查询
	esQuery := eq.ESQuery{
		// Query: eq.Bool(eq.WithMust([]eq.Map{eq.Match("name", "snow")})),
		Query: eq.Match("name", "snow"),
		Aggs:  eq.TermsAgg("name.keyword", eq.WithSize(8)),
	}
	fmt.Println(esQuery.JSON())

	l, t, err := eq.QueryList[Books](es, "books", esQuery)
	lj, _ := json.Marshal(l)
	fmt.Printf("%+v\n", string(lj))
	fmt.Println(t)

	raw, err := eq.QueryAgg[eq.TermsAggResult](es, "books", esQuery)
	fmt.Println(raw)
	l, t, err = QueryBooksByAuthorName(es, "Neal", "snow")
	JPrint(l)
	fmt.Println(t)

}

// QueryBooksByName 根据name查询books的详细数据
func QueryBooksByName(es *elasticsearch.Client, name string) ([]*Books, int, error) {
	esQuery := eq.ESQuery{
		Query: eq.Match("name", name),
	}

	l, t, err := eq.QueryList[Books](es, "books", esQuery)
	if err != nil {
		return nil, 0, err
	}
	return l, t, nil
}

// QueryBooksByAuthorName 根据author、name查询books的详细数据
func QueryBooksByAuthorName(es *elasticsearch.Client, author string, name string) ([]*Books, int, error) {
	queries := []eq.Map{
		eq.Match("author", author),
		eq.Match("name", name),
	}
	esQuery := eq.ESQuery{Query: eq.Bool(eq.WithMust(queries))}
	fmt.Println(esQuery.JSON())
	l, t, err := eq.QueryList[Books](es, "books", esQuery)
	if err != nil {
		return nil, 0, err
	}
	return l, t, nil
}
