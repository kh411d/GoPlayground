package main

import (
    "fmt"
    "reflect"
)

func DRYCreate(my interface{}) {

    v := reflect.ValueOf(my)
    myType := v.Elem()

    fmt.Printf("%#v\n\n", my)
    fmt.Printf("%#v\n\n", v)
    fmt.Printf("%#v\n\n", myType.Field(0))
    fmt.Printf("%#v\n\n", myType.Type().Field(0))

    fmt.Printf("%#v\n\n", v.Interface())
    fmt.Printf("%v\n\n", myType)
    fmt.Printf("%v\n\n", myType.Type())

    x := reflect.New(myType.Type())

    fmt.Printf("%#v\n\n", x)

    x.Elem().Set(myType)

}

type User struct {
    name string
}

func main() {
    DRYCreate(&User{name: "kambing"})
}
