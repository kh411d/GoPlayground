package main

import (
    "fmt"
)

type A struct {
    vara string
    varb int
}

func (t *A) check() string {

    return "on A"
}

//embed A to B
type B struct {
    A
}

func (t *B) onB() {
    fmt.Println(t.check())
}

func main() {
    //init embedded struct variable
    x := &B{A: A{vara: "kambing"}}
    //x.onB()
    fmt.Println(x.vara)
    fmt.Printf("%#v", x)
}
