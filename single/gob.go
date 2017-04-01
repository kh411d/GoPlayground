package main

import (
    "fmt"
    "encoding/gob"
    "bytes"
    "crypto/md5"
)

var m = map[string]int{"one":1, "two":2, "three":3}

func main() {
    b := new(bytes.Buffer)
    //b.Write([]byte("0eff81040102ff8200010c0104000015ff820003036f6e65020374776f0405746872656506"))
   
    fmt.Printf("%x",b)
    fmt.Println("")

    e := gob.NewEncoder(b)



    // Encoding the map
    err := e.Encode(m)
    if err != nil {
        panic(err)
    }

     s := md5.Sum(b)
    fmt.Printf("%x",string(s[:]))
    fmt.Println("")

    fmt.Printf("%x",b)
    fmt.Println("")

    var decodedMap map[string]int
    d := gob.NewDecoder(b)

    // Decoding the serialized data
    err = d.Decode(&decodedMap)
    if err != nil {
        panic(err)
    }

    // Ta da! It is a map!
    fmt.Printf("%#v\n", decodedMap)
}