package demo

import (
	"errors"
	"reflect"
	"testing"

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
			input: User{Name: "Jie", Sex: 1},
			wantRes: map[string]interface{}{
				"Name": "Jie",
				"Sex": 1,
				"age": 0,
			},
			wantErr: nil,
		},
		{
			name: "&user",
			input: &User{Name: "Shu", Sex: 0},
			wantRes: map[string]interface{}{
				"Name": "Shu",
				"Sex": 0,
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

