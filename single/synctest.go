package main

import (
    "fmt"
    "sync"
    "time"
)

type Stat struct {
    sync.Mutex
    counters map[string]int64
}

type StatAverage Stat

func (s *Stat) increaseCounter(key string, no int) {
    s.Lock()
    defer s.Unlock()
    fmt.Printf("%d\n", no)
    if c, exists := s.counters[key]; !exists {
        s.counters[key] = 1
        fmt.Println("not exists")
    } else {
        s.counters[key] = c + 1
        fmt.Println("exists")
    }
}

func main() {
    s := &Stat{
        counters: make(map[string]int64),
    }

    go s.increaseCounter("test", 1)
    go s.increaseCounter("test", 2)
    s.increaseCounter("test", 3)
    time.Sleep(time.Millisecond)
    fmt.Printf("%#v\n", s.counters)
}
