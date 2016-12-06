package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := Foo{Name: "name"}
	test.Bar = &Bar{&test}
	test.Test()
}

type Foo struct {
	*Bar
	Name string
}

func (s *Foo) Method() {
	fmt.Println("Foo.Method()")
}

type Bar struct {
	context interface{}
}

func (s *Bar) Test() {

	// Debug Bar composition
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	fmt.Printf("model: self:%#v\nt:%#v\nv:%#v\n", s, t, v)
	// List fields of s.context
	for i := 0; i < reflect.Indirect(reflect.ValueOf(s.context)).NumField(); i++ {
		fmt.Println("Field:", reflect.Indirect(reflect.ValueOf(s.context)).Field(i))
	}
	// List methods of s.context
	for i := 0; i < reflect.TypeOf(s.context).NumMethod(); i++ {
		fmt.Println("Method:", reflect.TypeOf(s.context).Method(i).Name)
	}
	// Test method call on top-level method
	if method := reflect.ValueOf(s.context).MethodByName("Method"); !method.IsNil() {
		method.Call(nil)
	}
}
