package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Z struct {
	Id int
}

type V struct {
	Id int
	F  Z
}

type T struct {
	Id int
	F  V
	X  []int
	K  I
}

type I interface {
	lompat() string
}

func InspectStructV(val reflect.Value) {

	fmt.Printf("\n\n")

	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		fmt.Println("dor")
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			fmt.Println("Gotcha is reflect Interface")
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		fmt.Printf("Val Ptr %v\n", val.Kind())
	}

	fmt.Printf("Val %#v\n", val)

	if _, ok := val.Type().FieldByName("KategoriKuda"); ok {
		x := val.MethodByName("Makan").Call([]reflect.Value{})
		fmt.Printf("%#v\n", x[0].Interface().(error))
	}

	fmt.Printf("Val Kind %v\n", val.Kind())
	fmt.Printf("Total field %d\n", val.NumField())

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		address := "not-addressable"

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()

		}
		if valueField.CanAddr() {
			address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
		}

		if valueField.Kind() == reflect.Slice {

			//fmt.Printf("Ini total slice %v \n", valueField.Len())
			//fmt.Printf("Ini tipe slice %v \n", typeField.Type)

			/*
			       var inint int
			   for i := 0; i < valueField.Len(); i++ {
			       fmt.Println(valueField.Index(i).Interface())
			       inint = valueField.Index(i).Interface().(int)
			       fmt.Println(inint)
			       fmt.Printf("Type %v \n", valueField.Index(i).Type())

			   }*/
		}

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n",
			typeField.Name,
			valueField.Interface(),
			address,
			typeField.Type,
			valueField.Kind())

		if valueField.Kind() == reflect.Struct {
			InspectStructV(valueField)
		}
	}
}

func InspectStruct(v interface{}) {

	//fmt.Printf("TypeOf Ptr %v\n", v.(kuda))
	fmt.Printf("TypeOf Ptr %v\n", reflect.TypeOf(v))
	fmt.Printf("TypeOf Ptr %v\n", reflect.ValueOf(v).Type())
	fmt.Printf("TypeOf Ptr %v\n", reflect.TypeOf(v).Elem())

	InspectStructV(reflect.ValueOf(v))
}

type KategoriKuda struct {
	Warna string
}
type kuda struct {
	Nama string
	KategoriKuda
}

func (KategoriKuda) Makan() error {
	return errors.New("kuda makan")
}

func (k kuda) lompat() string { return "lompat" }

func main() {
	/*t := new(T)
	  t.Id = 1
	  t.X = []int{1, 2, 3, 4, 4}
	  t.F = *new(V)
	  t.F.Id = 2
	  t.F.F = *new(Z)
	  t.F.F.Id = 3
	  t.K = new(kuda)*/
	// t := (*I)(nil)

	// t := (I)(new(kuda))
	// fmt.Printf("ini adalah %v\n", t)

	k := kuda{
		Nama: "Si Kuda",
	}
	k.Warna = "hitam"
	InspectStruct(&k)
}
