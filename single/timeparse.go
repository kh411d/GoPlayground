package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    loopDate := now.AddDate(0, 0, 1)
    t1, _ := time.Parse(time.RFC3339, loopDate.Format(`2006-01-02T`)+`15:00:00`+"+07:00")
    t2, _ := time.Parse(time.RFC3339, loopDate.Format(`2006-01-02T`)+`17:00:00`+"+07:00")
    fmt.Println(now)
    fmt.Println(loopDate)
    fmt.Println(t1)
    fmt.Println(t2)
    if now.After(t1) && now.Before(t2) { // In the delivery slot
        fmt.Println("asdf")
    } else if now.After(t2) { // After the delivery slot
        fmt.Println("qwer")
    }

}
