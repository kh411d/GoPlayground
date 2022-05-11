package main

import (
	"fmt"
	"runtime"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	return f.Name()
}

func getLocation(skip int) {
	//if you need to view multiple callers, use https://go.dev/doc/go1.9#callersframes for performance wise
	//this code below only for single caller
	if pc, _, _, ok := runtime.Caller(skip); ok {
		fmt.Println(runtime.FuncForPC(pc).Name())
	}
}

func getErr() {

	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	fmt.Println(f.Name())
}

func PublicFn() {
	getLocation(1)
}

func ItuFn() {
	PublicFn()
	getLocation(1)
}

func main() {
	PublicFn()
	ItuFn()
}
