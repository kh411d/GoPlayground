package main

import (
    "encoding/json"
    "fmt"
)

type structB struct {
    Ini string
    Itu int
}

func main() {
    x := &structB{}

    err := json.Unmarshal([]byte(`{"ini":"asf","itu":2}`), &x)
    fmt.Println(err)
    x.Ini = "hoho"

    fmt.Printf("%#v\n\n", x)
    fmt.Printf("%#v\n\n", *x)
}
