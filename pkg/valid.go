package pkg

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMsg(err error, obj interface{}) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
			}
		}
	}
	return err.Error()
}
