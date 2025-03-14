package esquery

import "fmt"

/******** 动态计算时间 *******/

// NYearAgo 当前时间近几年内
func NYearAgo(n int) string {
	return fmt.Sprintf("now-%dy", n)
}

// NYearAgoWee 近几年 (整年对齐)
func NYearAgoWee(n int) string {
	return fmt.Sprintf("now-%dy/y", n)
}

// NQuarterAgo 当前时间近几季度内
func NQuarterAgo(n int) string {
	return fmt.Sprintf("now-%dq", n)
}

// NQuarterAgoWee 近几季度 (整季度对齐)
func NQuarterAgoWee(n int) string {
	return fmt.Sprintf("now-%dq/q", n)
}

// NMonthAgo 当前时间近几月内
func NMonthAgo(n int) string {
	return fmt.Sprintf("now-%dM", n)
}

// NMonthAgoWee 近几个月 (整月对齐)
func NMonthAgoWee(n int) string {
	return fmt.Sprintf("now-%dM/M", n)
}

// NWeekAgo 当前时间近几周内
func NWeekAgo(n int) string {
	return fmt.Sprintf("now-%dw", n)
}

// NWeekAgoWee 近几周 (整周对齐)
func NWeekAgoWee(n int) string {
	return fmt.Sprintf("now-%dw/w", n)
}

// NDayAgo 当前时间近几天内
func NDayAgo(n int) string {
	return fmt.Sprintf("now-%dd", n)
}

// NDayAgoWee 近几天 (整天对齐)
func NDayAgoWee(n int) string {
	return fmt.Sprintf("now-%dd/d", n)
}

// NHourAgo 当前时间近几小时内
func NHourAgo(n int) string {
	return fmt.Sprintf("now-%dh", n)
}

// NHourAgoWee 近几小时 (整点对齐)
func NHourAgoWee(n int) string {
	return fmt.Sprintf("now-%dh/h", n)
}

// NMinuteAgo 当前时间近几分钟内
func NMinuteAgo(n int) string {
	return fmt.Sprintf("now-%dm", n)
}

// NMinuteAgoWee 近几分钟 (整分对齐)
func NMinuteAgoWee(n int) string {
	return fmt.Sprintf("now-%dm/m", n)
}

// NSecondAgo 当前时间近几秒内
func NSecondAgo(n int) string {
	return fmt.Sprintf("now-%ds", n)
}

// NSecondAgoWee 近几秒 (整秒对齐)
func NSecondAgoWee(n int) string {
	return fmt.Sprintf("now-%ds/s", n)
}

// Now 获取当前时间
func Now() string {
	return "now"
}

// NowAligned 获取当前时间的整点 (按日对齐)
func NowAligned() string {
	return "now/d"
}

// Today 获取当天的起始时间
func Today() string {
	return "now/d"
}

// ThisWeek 获取本周的起始时间
func ThisWeek() string {
	return "now/w"
}

// ThisMonth 获取本月的起始时间
func ThisMonth() string {
	return "now/M"
}

// ThisQuarter 获取本季度的起始时间
func ThisQuarter() string {
	return "now/q"
}

// ThisYear 获取本年的起始时间
func ThisYear() string {
	return "now/y"
}
