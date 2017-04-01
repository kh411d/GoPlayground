package main

import (
        "fmt"
        "math"
        "reflect"
)

type Circle struct {
        r float64
}

func (c *Circle) Area() float64 {
        return math.Pi * c.r * c.r
}

func main() {
        c := Circle{1.2}

        // call Area() method
        fmt.Println(c.Area())

        // call Area() method by Name
        v := reflect.ValueOf(&c).MethodByName("Area").Call([]reflect.Value{})
        fmt.Println(v[0].Float())
}
