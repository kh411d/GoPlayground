package main

import "fmt"

func main() {
    // create new channel of type int
    ch := make(chan int, 3)

    // start new anonymous goroutine
    go func() {
        // send 42 to channel
        fmt.Println("send 42")
        ch <- 42
        fmt.Println("send 412")
        ch <- 412
        fmt.Println("send 4")
        ch <- 4
        close(ch)
    }()
    // read from channel
    for v := range ch {
        fmt.Println(v)
    }

}
