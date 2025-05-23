package register

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// makeParam 尝试修复并构造函数参数
func makeParam(datas []string, parmas [][]string) ([]reflect.Value, error) {
	argVals := []reflect.Value{}
	for idx, pm := range parmas {
		if len(datas) > idx {
			data, err := Str2type(datas[idx], pm[1])
			if err != nil {
				return nil, fmt.Errorf("param type error")
			}
			argVals = append(argVals, reflect.ValueOf(data))
		} else {
			data := Zero(pm[1], pm[2])
			argVals = append(argVals, reflect.ValueOf(data))
		}
	}
	return argVals, nil
}

// paramConv 参数转换
func paramConv(query string, datas []string, parmas [][]string) ([]reflect.Value, error) {
	if !checkParam(datas, parmas) {
		return makeParam(datas, parmas)
	}

	var optType string
	argVals := make([]reflect.Value, 0, len(datas))
	for idx, data := range datas {
		var paramType string
		if len(optType) > 0 {
			paramType = optType
		} else {
			paramType = parmas[idx][1]
			if strings.HasPrefix(paramType, "...") {
				optType = paramType
			}
		}

		pm, err := Str2type(data, paramType)
		if err != nil {
			return nil, err
		}
		argVals = append(argVals, reflect.ValueOf(pm))
	}

	return argVals, nil
}

// Str2type 将字符串内容转换成对应的类型数据
func Str2type(data, paramType string) (any, error) {
	switch paramType {
	case "int", "...int":
		i, err := strconv.Atoi(data)
		if err != nil {
			return int(0), nil
		}
		return int(i), nil
	case "int32", "...int32":
		i, err := strconv.ParseInt(data, 10, 32)
		if err != nil {
			return int32(0), nil
		}
		return int32(i), nil
	case "int64", "...int64":
		i, err := strconv.ParseInt(data, 10, 32)
		if err != nil {
			return int64(0), nil
		}
		return int64(i), nil
	case "string", "...string":
		return data, nil
	case "float32", "...float32":
		i, err := strconv.ParseFloat(data, 10)
		if err != nil {
			return float32(0), nil
		}
		return i, nil
	case "float64", "...float64":
		i, err := strconv.ParseFloat(data, 10)
		if err != nil {
			return float64(0), nil
		}
		return i, nil
	case "time.Time", "...time.Time":
		t, err := time.Parse("2006-01-02 15:04:05", data)
		if err != nil {
			t = startOfDay(time.Now(), 0)
		}
		return t, nil
	}
	return nil, fmt.Errorf("unknown type:%s", paramType)

}

// OptParamTrim 去掉可选参数符号
func OptParamTrim(s string) string {
	return strings.Trim(s, "...")
}

// checkParam 检查参数个数是否满足
func checkParam(datas []string, parmas [][]string) bool {
	var must, opt int
	for _, p := range parmas {
		if len(p) >= 2 {
			if strings.HasPrefix(p[1], "...") {
				opt++
			} else {
				must++
			}
		}
	}

	if len(datas) == len(parmas) ||
		(len(datas) >= must && len(datas) <= must+opt) {
		return true
	}

	return false
}

// 日期的起始时间
func startOfDay(t time.Time, n int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()-n, 0, 0, 0, 0, t.Location())
}

// Zero 构造类型零值
func Zero(typ, comment string) any {
	var n int = 5     // 最近的默认值为5
	var f float32 = 5 // 浮点数默认值为5
	switch typ {
	case "int", "...int":
		if strings.Contains(comment, "近") || strings.Contains(comment, "几") {
			return int(n)
		}
		return int(0)
	case "int32", "...int32":
		if strings.Contains(comment, "近") || strings.Contains(comment, "几") {
			return int32(n)
		}
		return int32(0)
	case "int64", "...int64":
		if strings.Contains(comment, "近") || strings.Contains(comment, "几") {
			return int64(n)
		}
		return int64(0)
	case "string", "...string":
		return ""
	case "float32", "...float32":
		if strings.Contains(comment, "近") || strings.Contains(comment, "几") {
			return float32(f)
		}
		return float32(0)
	case "float64", "...float64":
		if strings.Contains(comment, "近") || strings.Contains(comment, "几") {
			return float64(f)
		}
		return float64(0)
	case "time.Time", "...time.Time":
		if strings.Contains(comment, "近") || strings.Contains(comment, "几") {
			return startOfDay(time.Now(), n)
		}
		return startOfDay(time.Now(), 0)
	}
	return nil
}
