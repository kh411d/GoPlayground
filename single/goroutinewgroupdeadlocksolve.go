package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    done := make(chan string)
    wq := make(chan interface{})
    workerCount := 2

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go doit(i, wq, done, &wg)
    }

    for i := 0; i < workerCount; i++ {
        wq <- i
    }

    close(done)

    wg.Wait()
    fmt.Println("all done!")
}

func doit(workerId int, wq <-chan interface{}, done <-chan string, wg *sync.WaitGroup) {
    fmt.Printf("[%v] is running\n", workerId)
    defer wg.Done()
    for {
        select {
        case m := <-wq:
            fmt.Printf("[%v] m => %v\n", workerId, m)
        case x := <-done:
            fmt.Printf("[%v] is done %s \n", workerId, x)
            return
        }
    }
}
