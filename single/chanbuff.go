package main

import (
    "fmt"
    "strconv"
)

func main() {

    message := make(chan string)

    go func(message chan string) {
        for i := 1; i <= 5; i++ {
            fmt.Println("write ", i)
            message <- "message " + strconv.Itoa(i)
        }
    }(message)

    go func(message chan string) {
        for i := 1; i <= 5; i++ {
            select {
            case x := <-message:
                fmt.Println("read ", x)
            }
        }
    }(message)

    fmt.Scanln()

}
