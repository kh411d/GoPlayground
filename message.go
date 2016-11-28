 package main

import (
     "fmt"
    // "time"
)
func main() {
     /*message := make(chan string,2) // no buffer
     count := 5

     go func() {
          for i := 1; i <= count; i++ {
               fmt.Println("send message")
               message <- fmt.Sprintf("message %d", i)
          }
     }()

     time.Sleep(time.Second * 3)

     for i := 1; i <= count; i++ {
          fmt.Println(<-message)
     }*/

      /*message := make(chan string)
     count := 3

     go func() {
          for i := 1; i <= count; i++ {
               message <- fmt.Sprintf("message %d", i)
          }
          close(message)
     }()

     for msg := range message {
          fmt.Println(msg)
     }*/

      done := make(chan bool)

     go func() {
          println("goroutine message")

          //done <- true
          // Just send a signal "I'm done"
          close(done)
     }()

     println("main function message")
     fmt.Println(<-done)
     <-done
}
 