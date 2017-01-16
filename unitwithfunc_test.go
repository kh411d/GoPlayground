package unitwithfunc

import (
    "fmt"
    "testing"
)

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

func Thefunc(a A_interface) {
    fmt.Println("running the func")
    a.method1()
    a.method2()
}

/*package unitwithfunc_test

import (
    "fmt"
    "testing"
)*/

type mock struct{}

func (self *mock) method1() { fmt.Println("test method 1 execute") }
func (self *mock) method2() { fmt.Println("test method 2 execute") }

func TestThefunc(t *testing.T) {
    x := &mock{}
    Thefunc(x)
}
