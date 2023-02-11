package demo

import (
	"errors"
	"fmt"
	"reflect"
)

// IterateField 用反射输出字段名字和值
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
			// 不公开字段，reflect 能拿到类型信息，但是拿不到值；
			// 硬要处理的话，零值填充，反而多此一举。
			res[fd.Name] = reflect.Zero(fd.Type).Interface()
		}
		
	}
	return res, nil
}



// setField 用反射设置值
func SetField(entity interface{}, field string, newVal interface{}) {
	if err := setField(entity, field, newVal); err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Printf("%#v\n", entity)
}

func setField(entity interface{}, field string, newVal interface{}) error {
	val := reflect.ValueOf(entity)
	typ := val.Type()

	// 只能是一级指针，类似 *User
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}	

	typ = typ.Elem()
	val = val.Elem()
	
	// 这个地方判断不出来 field 在不在
	fd := val.FieldByName(field)

	// 利用 type 来判断 field 在不在
	if _, found := typ.FieldByName(field); !found {
		return errors.New("字段不存在")
	}

	if !fd.CanSet() {
		return errors.New("字段不可修改")
	}

	fd.Set(reflect.ValueOf(newVal))
	return nil
}