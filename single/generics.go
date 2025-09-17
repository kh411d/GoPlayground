package main

import (
	"fmt"
	"reflect"
)

// type alphanum interface {
// 	int | int8 | int16 | int32 | int64 | string
// }

// type DupCount[T comparable] map[T]int

// func (d DupCount[T]) Add(key T) {
// 	d[key]++
// }

// func (d DupCount[T]) AnyDup() (isdup bool) {
// 	for _, v := range d {
// 		if v > 1 {
// 			isdup = true
// 			break
// 		}
// 	}
// 	return
// }

type DupCount map[interface{}]int

func (d DupCount) Add(keyI interface{}) {
	if key, ok := keyI.(comparable); ok {
		d[key]++
	}
	// if d.isAllowed(key) {
	// 	d[key]++
	// }
}

func (d DupCount) AnyDup() (isdup bool) {
	for _, v := range d {
		if v > 1 {
			isdup = true
			break
		}
	}
	return
}

func (d DupCount) isAllowed(key interface{}) (ok bool) {
	allowedType := []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String,
	}

	for _, v := range allowedType {
		if v == reflect.ValueOf(key).Kind() {
			ok = true
			break
		}
	}
	return
}

func main() {

	x := make(DupCount)
	x.Add(23424)
	x.Add("asdfasdf")

	// x.Add("kabal")
	// x.Add("kaball")
	//x.Add("asdfasdf")
	fmt.Println(x.AnyDup())
}
