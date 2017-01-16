package main

import "fmt"

type A_interface interface {
    method1()
    method2()
}

type methods struct{}

func (self *methods) method1() {
    fmt.Println("method 1 execute")
}

func (self *methods) method2() {
    fmt.Println("method 2 execute")
}

func (self *methods) method3() {
    fmt.Println("method 3 execute")
}

func Thefunc(a A_interface) {
    fmt.Println("running the func")
    a.method1()
    a.method2()
}

type coba1 struct{}

func (self *coba1) act() { fmt.Println("act coba1") }

type cobaparent struct {
    item string
    coba1
}

func main() {
    a := &methods{}

    /*var x A_interface
      x = a
      x.method3()*/
    c := &cobaparent{}
    c.act()
    Thefunc(a)
}
