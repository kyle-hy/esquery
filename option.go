package esquery

// Option ES 各种参数属性
type Option func(Map)

// newOptMap 构造属性
func newOptMap(opts []Option) Map {
	m := make(Map)
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithBoost 提升某个字段在相关度评分中的优先级
// @param value 权重值
func WithBoost(value any) Option {
	return func(m Map) {
		m["boost"] = value
	}
}

// WithMinimumShouldMatch 控制查询匹配的文档必须包含多少个查询值
// @param value 个数或百分比
func WithMinimumShouldMatch(value any) Option {
	return func(m Map) {
		m["minimum_should_match"] = value
	}
}

// operator的枚举值
var (
	AND = "AND"
	OR  = "OR"
)

// WithOperator 词与词之间的逻辑关系 (AND/OR)
// @param value AND/OR(默认)
func WithOperator(value any) Option {
	return func(m Map) {
		m["operator"] = value
	}
}

// fuzziness的枚举值
var (
	AUTO = "AUTO" // 根据查询词长度自动调整模糊度(推荐)
)

// WithFuzziness 允许的最大拼写错误距离 (模糊匹配)
// @param value AUTO/个数
func WithFuzziness(value any) Option {
	return func(m Map) {
		m["fuzziness"] = value
	}
}

// fuzziness的枚举值
var (
	Standard   = "standard"   // 标注分词器(默认)
	Whitespace = "whitespace" // 空白符分词
)

// WithAnalyzer 指定分词器，替代默认分词器
// @param value 分词器类型standard/whitespace
func WithAnalyzer(value string) Option {
	return func(m Map) {
		m["analyzer"] = value
	}
}

// zero_terms_query的枚举值
var (
	NONE = "none" // 返回空
	ALL  = "all"  // 全部返回
)

// WithZeroTermsQuery 当查询为空时的行为
// @param value 空查询时返回方式none/all
func WithZeroTermsQuery(value string) Option {
	return func(m Map) {
		m["zero_terms_query"] = value
	}
}

// type匹配类型枚举值
var (
	BestFields   = "best_fields"   // 适用于单个最佳字段匹配，侧重相关性高的字段。最适合典型的全文搜索
	MostFields   = "most_fields"   // 在多个字段中查找并合并评分，适用于字段可能有重复信息（如标题、副标题）
	CrossFields  = "cross_fields"  // 将查询词按单词拆分，跨字段匹配，适用于人名、地址等数据
	Phrase       = "phrase"        // 严格遵守短语匹配顺序
	PhrasePrefix = "phrase_prefix" // 执行短语前缀匹配，允许部分前缀匹配
)

// WithType 匹配类型，控制多字段搜索的策略
// @param value 匹配类型
func WithType(value string) Option {
	return func(m Map) {
		m["type"] = value
	}
}

// WithFormat 日期格式
// @param value 日期格式（如 "yyyy-MM-dd"）
func WithFormat(value string) Option {
	return func(m Map) {
		m["format"] = value
	}
}
