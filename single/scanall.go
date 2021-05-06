package main

import (
	"errors"
	"fmt"
	"reflect"
)

func test(outputVal interface{}) (err error) {
	value := reflect.ValueOf(outputVal)

	if value.Kind() != reflect.Ptr {
		err = errors.New("not pointer")
		return
	}

	direct := reflect.Indirect(value)
	fmt.Printf("reflect indirect value: %#v\n", direct)

	slice := Deref(value.Type())
	if slice.Kind() != reflect.Slice {
		err = errors.New("not slice")
		return
	}

	fmt.Printf("ini slice elem: %s\n", slice.Elem())

	isPtr := slice.Elem().Kind() == reflect.Ptr
	base := Deref(slice.Elem())

	for i := 1; i <= 10; i++ {
		vp := reflect.New(base)

		//dst := vp.Interface()
		//dst[0]

		if isPtr {
			direct.Set(reflect.Append(direct, vp))
		} else {
			direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
		}
	}
	return
}

func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

type strc struct {
	Name string
}

func main() {

	//var x []strc
	var x []map[string]interface{}
	err := test(&x)

	fmt.Println(err)
	fmt.Printf("%#v", x)
}
