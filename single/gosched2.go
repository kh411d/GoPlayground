package main

import (
  "fmt"
  "runtime"
  "time"
)

const testBytes = `{ "Test": "value" }`

type Message struct {
  Test string
}

/*func cpuIntensive(p *Message) {
  for i := int64(1); i <= 1000; i++ {
    json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
    //runtime.Gosched()
  }
  fmt.Println("Done intensive thing")

}*/

func cpuIntensive(p *int64) {
  for i := int64(1); i <= 10000000; i++ {
    *p = i
  }
  fmt.Println("Done intensive thing")
}

func printVar(p *Message) {
  fmt.Printf("print x = %v\n", *p)
}

func main() {
  runtime.GOMAXPROCS(1)

  x := Message{}
  go cpuIntensive(&x)
  go printVar(&x)
  time.Sleep(1 * time.Nanosecond)
  fmt.Scanln()
}
