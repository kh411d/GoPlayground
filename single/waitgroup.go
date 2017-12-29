package main

import (
    "fmt"
    "log"
    "os"
    "sync"
)

var wg sync.WaitGroup

var fileMutex sync.Mutex

func test(s string, fo *os.File) {
    defer wg.Done()
    for i := 0; i < 105; i++ {
        fileMutex.Lock()
        fmt.Fprintf(fo, "%s%d", s, i)
        fileMutex.Unlock()
    }
}

func main() {
    fo, err := os.Create("output.txt")
    if err != nil {
        log.Fatal(err)
    }
    for i := 0; i < 4; i++ {
        wg.Add(1)
        go test("bye", fo)

    }
    wg.Wait()
}
