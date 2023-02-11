package demo

import (
	"errors"
	"fmt"
	"reflect"
)

// IterateField 用反射输出所有字段名字和值
func IterateFields(val interface{}) {
	res, err := iterateFields(val)
	if err != nil {
		fmt.Println(err)
		return 
	}

	for k, v := range res {
		fmt.Println(k, v)
	}
}


// 没有返回值，拆分，方便测试
func iterateFields(input interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, errors.New("不能为 nil")
	}


	typ :=reflect.TypeOf(input)
	val := reflect.ValueOf(input)

	// 处理指针，要拿到指针指向的东西
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	// 如果不是 struct，就返回 error
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("非法类型")
	}

	// 字段数量
	num := typ.NumField()
	res := make(map[string]interface{}, num)

	for i := 0; i < num; i++ {
		// 返回 struct 第 i 个`字段`类型信息和值信息
		fd := typ.Field(i)
		fdVal := val.Field(i)

		if fd.IsExported() {
			res[fd.Name] = fdVal.Interface() 
		} else {
			// 不公开字段，我们使用零值填充
			res[fd.Name] = reflect.Zero(fd.Type).Interface()
		}
		
	}
	return res, nil
}
