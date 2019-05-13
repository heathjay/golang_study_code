package test

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		sv    reflect.Type
		elem  reflect.Value
	)
	sv = reflect.TypeOf(model)                  //指针传入，变成reflect.Value
	t.Log("reflect.TypeOf", sv.Kind().String()) //ptr
	sv = sv.Elem()                              //sv指向的类型
	t.Log("reflect.TypeOf.Elem:", sv.Kind().String())

	elem = reflect.New(sv) //New返回一个Value类型，该值只有一个指向类型typ的新申请的零值的指针
	t.Log("reflect.new,", elem.Kind().String())
	t.Log("reflect.new.Elem(),", elem.Elem().Kind().String())

	model = elem.Interface().(*user) //*user指向elem同一块内容
	elem = elem.Elem()               //取得elem所指向的值
	elem.FieldByName("UserId").SetString("123456")
	elem.FieldByName("Name").SetString("nic")
	t.Log("model model.Name", model, model.Name)

}
