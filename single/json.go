package main

import (
    "fmt"
    "encoding/json"
    "crypto/md5"
)

var m = map[string]int{"one":1, "twjkkjo":2, "three":3}

func main() {
   
   x,_ := json.Marshal(m)

     s := md5.Sum(x)
    fmt.Printf("%x",string(s[:]))
    fmt.Println("")

    
}