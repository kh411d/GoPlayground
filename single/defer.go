package main

import "fmt"

type xxx struct {
    itu string
}

func test() string {

    x := &xxx{
        itu: "ini dia",
    }

    defer func(x *xxx) string {
        x.itu = "kambing"
        return x.itu
    }(x)

    return x.itu
}

func main() {
    fmt.Println(test())
}
