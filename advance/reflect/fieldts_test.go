package demo

import (
	"errors"
	"reflect"
	"testing"
	"github.com/sjxiang/one/advance/reflect/types"

	"github.com/stretchr/testify/assert"
)


func TestReflectPanic(t *testing.T) {
	typ := reflect.TypeOf(User{})  // 换 &User{}，测试结果 panic

	if typ.Kind() == reflect.Struct {
		t.Log("结构体")
	} else if typ.Kind() == reflect.Ptr {
		t.Fatalf("指针，失策了")
	}

	t.Logf("User{} 字段数量：%d\n", typ.NumField())
}


type User struct {
	Name string
	Sex  int
	age  int  // 不公开字段，reflect 能拿到类型信息，但是拿不到值；硬要处理的话，零值填充，反而多此一举。
}




// TDD 测试驱动开发
func TestIterateFields(t *testing.T) {
	
	tests := []struct {
		name    string
		// 输入部分
		input     interface{}
		// 输出部分
		wantRes map[string]interface{}
		wantErr error
	}{	
		{
			name: "nil",
			input: nil,
			wantErr: errors.New("不能为 nil"),
		},
		{   
			name: "user",
			input: &types.User{Name: "Jie"},
			wantRes: map[string]interface{}{
				"Name": "Jie",
				"age": 0,
			},
			wantErr: nil,
		},
		{
			name: "&user",
			input: &types.User{Name: "Jie"},
			wantRes: map[string]interface{}{
				"Name": "Jie",
				"age": 0,
			},
			// 考虑支持
			wantErr: nil,
		},
		{
			name: "slice",
			input: []string{},
			wantErr: errors.New("非法类型"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := iterateFields(tt.input)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return 
			}

			assert.Equal(t, tt.wantRes, res)
		})
	}
}


func TestSetField(t *testing.T) {

	tests := []struct {
		name    string
		// 输入部分
		entity     interface{}
		field   string
		newVal     interface{}
		// 输出部分
		wantErr error
	}{	
		{
			name: "struct",
			entity: types.User{},
			wantErr: errors.New("非法类型"),
		},
		{
			name: "invalid field",
			entity: &types.User{},
			field: "age",
			newVal: "25",
			wantErr: errors.New("字段不可修改"),
		},
		{
			name: "Not available field",
			entity: &types.User{},
			field: "Sex",
			newVal: "1",
			wantErr: errors.New("字段不存在"),
		},
		{
			name: "pass",
			entity: &types.User{Name: "姜云升"},
			field: "Name",
			newVal: "法老",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := setField(tt.entity, tt.field, tt.newVal)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

