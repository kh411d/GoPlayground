package main

import "fmt"

func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {

    /** 1st Method *
      // Set up the pipeline.
      c := gen(2, 3)
      out := sq(c)
      // Consume the output.
      fmt.Println(<-out) // 4
      fmt.Println(<-out) // 9
    /**/

    /** 2nd Method **
      // Set up the pipeline and consume the output.
      for n := range sq(gen(2, 3)) {
          fmt.Println(n) // 16 then 81
      }
      /**/

    /** 3rd Method **/
    // Set up the pipeline and consume the output.
    c := gen(2, 3)
    out := sq(c)
    for r := 0; r < 2; r++ {
        select {
        case msg := <-out:
            fmt.Println(msg)
        }
    }
    /**/
}
