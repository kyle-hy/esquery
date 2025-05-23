package register

import (
	"fmt"
	"reflect"
)

// Condition 查询条件
type Condition map[string]any

// Data 返回的数据
type Data map[string]any

// APIFormat api函数模板格式
type APIFormat func(...any) (Condition, Data, error)

// 全量的函数
var handlers = map[string]*FuncInfo{}

// FuncInfo 函数信息
type FuncInfo struct {
	Func    reflect.Value // 函数反射的值
	Name    string        // 函数名
	Comment string        // 函数注释
	Params  [][]string    // 参数信息列表
}

// GetFunc 查询函数信息
func GetFunc(name string) (f *FuncInfo, ok bool) {
	f, ok = handlers[name]
	return
}

// SetFunc 查询函数信息，无锁，仅限init函数使用
func SetFunc(fi *FuncInfo) (f *FuncInfo, ok bool) {
	handlers[fi.Name] = fi
	return
}

// Handle 处理请求
// query 问题
// ps API接口及参数值列表
func Handle(query string, ps []string) (any, error) {
	if len(ps) == 0 {
		return nil, fmt.Errorf("empty param")
	}

	// API接口名称
	f, ok := handlers[ps[0]]
	if !ok {
		return nil, fmt.Errorf("unsupport %s", ps[0])
	}

	// 根据API参数说明转换入参格式
	args, err := paramConv(query, ps[1:], f.Params)
	if err != nil {
		return nil, err
	}

	out := f.Func.Call(args)
	if len(out) >= 3 && out[2].Interface() != nil {
		return nil, out[2].Interface().(error)
	}

	result := map[string]any{
		"api":     f.Name,
		"comment": f.Comment,
		"detail":  Condition{"cond": out[0].Interface(), "data": out[1].Interface()},
	}
	return result, nil
}
