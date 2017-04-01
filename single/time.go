package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    then := now.Add(-179 * time.Second)
    fmt.Println(now)
    fmt.Println(then)
}
