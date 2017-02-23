package main

import (
    "fmt"
    "reflect"
)

type Interface2 interface {
    GetName() string
}

func InterfaceOf(ifacePtr interface{}) reflect.Type {
    t := reflect.TypeOf(ifacePtr)
    for t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() != reflect.Interface {
        fmt.Println("Called InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
    }
    return t
}

// FromPtrTypeOf is to get real type from a pointer to value
func FromPtrTypeOf(obj interface{}) reflect.Type {
    realType := reflect.TypeOf(obj)
    for realType.Kind() == reflect.Ptr {
        realType = realType.Elem()
    }
    return realType
}

// FromPtrType is to get real type from a pointer to type
func FromPtrType(typ reflect.Type) reflect.Type {
    realType := typ
    for realType.Kind() == reflect.Ptr {
        realType = realType.Elem()
    }
    return realType
}

func main() {
    x := (*Interface2)(nil)
    c := InterfaceOf(x)
    fmt.Println(c)

    d := FromPtrType(c)
    fmt.Println(d)
}
