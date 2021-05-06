package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import "fmt"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) string {
    // write your code in Go 1.4
    var temp string
    var i int
    for i = 1; i <= N; i++ {
        if i%2 == 0 {
            temp = fmt.Sprintf(temp + "+")
        } else {
            temp = fmt.Sprintf(temp + "-")
        }
    }
    return temp
}

func main() {

}
