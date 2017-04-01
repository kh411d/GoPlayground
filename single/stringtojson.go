package main

import (
    "encoding/json"
    "fmt"
)

type Response2 struct {
    Barangnya []string `json:"barangnya"`
}

func main() {
    fmt.Println("Hello, playground")

    str := `{"barangnya":["asdfasdf","asdfasdfa"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Printf("%#v", res)

}
