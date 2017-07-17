package main

import (
     "fmt"
     "strconv"
     "time"
)

func runner(idx int, c chan []int, bucket *ember) {
     var temp []int
     time.Sleep(time.Millisecond * 1000)
     for k, v := range bucket.isi {
          if v != 0 {
               temp = append(temp, v)
               bucket.isi[k] = 0
          }
          if len(temp) == 5 {
               break
          }
     }
     fmt.Println("runner " + string(idx))
     fmt.Println(bucket)

     c <- temp

}

func runner2(idx int, c chan []int, b chan []int) {
     var temp []int
     fmt.Println("runner " + strconv.Itoa(idx))
     bucket := <-b
     for k, v := range bucket {
          if v != 0 {
               temp = append(temp, v)
               bucket[k] = 0
          }
          if len(temp) == 5 {
               break
          }
     }

     //Strange condition if channel c is sent after channel b is sent,
     //somehow the last runner is not being proceed
     b <- bucket
     c <- temp

}

func printer(c chan []int) {
     for {
          // msg := <-c
          // fmt.Println(msg)
          // time.Sleep(time.Second * 1)

          select {
          case msg := <-c:
               fmt.Println(msg)
               time.Sleep(time.Second * 1)
          }

     }
}

type ember struct {
     isi []int
}

func main() {

     data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

     //bucket := &ember{isi: data}

     c := make(chan []int, 5)
     bucket := make(chan []int)

     go runner2(1, c, bucket)
     go runner2(2, c, bucket)
     go runner2(3, c, bucket)
     go runner2(4, c, bucket)

     bucket <- data

     go printer(c)

     var input string
     fmt.Scanln(&input)

}
