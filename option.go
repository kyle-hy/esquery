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

// format 日期格式少量枚举
var (
	Format1      = "yyyy-MM-dd"          // "2024-03-14"
	Format2      = "yyyy/MM/dd"          // "2024/03/14"
	Format3      = "yyyy-MM-dd HH:mm:ss" // "2024-03-14 23:59:59"
	Format4      = "yyyy/MM/dd HH:mm:ss" // "2024/03/14 23:59:59"
	FormatSecond = "epoch_second"        // Unix 时间戳（秒）
	FormatMillis = "epoch_millis"        // Unix 时间戳（毫秒）
)

// WithFormat 日期格式
// @param value 日期格式（如 "yyyy-MM-dd"）
func WithFormat(value string) Option {
	return func(m Map) {
		m["format"] = value
	}
}

// 控制评分模式
var (
	AVG = "avg" //（默认）：取所有匹配文档得分的平均值。
	MIN = "min" // 取最低得分。
	MAX = "max" // 取最高得分。
	SUM = "sum" // 取所有匹配文档得分的总和。
	// NONE = "none" // 不计算得分。
)

// WithScoreMode 评分模式
// @param value 日期格式（如 "yyyy-MM-dd"）
func WithScoreMode(value string) Option {
	return func(m Map) {
		m["score_mode"] = value
	}
}

// WithParams 传递参数
// @param key 参数名
// @param value 参数值
func WithParams(key string, value any) Option {
	return func(m Map) {
		m[key] = value
	}
}

// WithCaseInsensitive 大小写敏感
// @param value 是否大小写敏感，默认false
func WithCaseInsensitive(value bool) Option {
	return func(m Map) {
		m["case_insensitive"] = value
	}
}

// WithIgnoreUnmapped 未映射字段处理
// @param value 是否忽略，默认false
func WithIgnoreUnmapped(value bool) Option {
	return func(m Map) {
		m["ignore_unmapped"] = value
	}
}

// 验证地理坐标方式
var (
	ARC   = "arc"   // 默认,使用球面几何计算距离，精度高但性能较低
	Plane = "plane" // 平面几何计算距离，性能较高但精度较低（适用于小范围）
)

// WithDistanceType 计算距离的方式
// @param value 是否忽略，默认false
func WithDistanceType(value string) Option {
	return func(m Map) {
		m["distance_type"] = value
	}
}

// 验证地理坐标方式
var (
	Strict          = "STRICT"           // 严格验证，无效坐标会抛出异常
	IgnoreMalformed = "IGNORE_MALFORMED" // 忽略无效坐标
	Coerce          = "COERCE"           // 尝试修正无效坐标
)

// WithValidationMethod 指定如何验证地理坐标
// @param value 是否忽略，默认false
func WithValidationMethod(value string) Option {
	return func(m Map) {
		m["validation_method"] = value
	}
}
