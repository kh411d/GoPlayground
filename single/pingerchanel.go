package main

import (
  "fmt"
  "time"
)

func pinger(c chan<- string) { // c can only be sent to
  //func pinger(c chan string) {
  for {
    c <- "ping"
  }
}

func ponger(c chan<- string) { // c can only be sent to
  //func ponger(c chan string) {
  for {
    c <- "pong"
  }
}

func faker(c chan<- string) { // c can only be sent to
  //func ponger(c chan string) {
  for {
    c <- "faker"
  }
}

func printer(c <-chan string, f <-chan string) { // c can only be receive to
  //func printer(c chan string) {
  for {

    // either example 1 or 2 has the same result

    // example 1: One channel only
    // msg := <-c
    // fmt.Println(msg)
    // time.Sleep(time.Second * 1)

    // example 2: Can be use for lot channel
    select {
    case msg := <-c:
      fmt.Println(msg)
      time.Sleep(time.Second * 1)
    //case msg := <-f:
    //fmt.Println(msg)
    //time.Sleep(time.Second * 1)
    case <-time.After(time.Second):
      fmt.Println("timeout")
    }
  }
}

func main() {
  var c chan string = make(chan string)
  //if buffer is set , all the buffer will be used first by first go routine,
  //as example: it will print pong, pong, pong, pong, pong and then pong, ping, pong, ping, ....
  //var c chan string = make(chan string, 5)
  var f chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go faker(f)
  go printer(c, f)

  var input string
  fmt.Scanln(&input)
}
