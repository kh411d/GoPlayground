package main

import (
    "fmt"
    "os"
    "runtime"
    "strconv"
    "sync"
)

func test(s string, fo *os.File) {
    var s1 [105]int
    count := 0
    for x := 1000; x < 1101; x++ {
        s1[count] = x
        count++
    }

    //fmt.Println(s1[0])
    for i := range s1 {
        runtime.Gosched()
        sd := s + strconv.Itoa(i)
        var fileMutex sync.Mutex
        fileMutex.Lock()
        fmt.Fprintf(fo, sd)
        defer fileMutex.Unlock()
    }
}

func main() {
    fo, err := os.Create("outputwrong.txt")
    if err != nil {
        panic(err)
    }
    for i := 0; i < 4; i++ {
        go test("bye", fo)

    }

}
