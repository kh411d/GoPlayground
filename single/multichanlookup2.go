package main

import (
     "fmt"
     "strconv"
)

const (
     PartSlice = 5
)

func runner(idx int, c chan []int, b chan []int) {
     var temp []int
     fmt.Println("runner " + strconv.Itoa(idx))
     bucket := <-b
     for k, v := range bucket {
          if v != 0 {
               temp = append(temp, v)
               bucket[k] = 0
          }
          if len(temp) == PartSlice {
               break
          }
     }

     //Strange condition if channel c is sent after channel b is sent,
     //somehow the last chan is not being sent
     b <- bucket
     c <- temp

}

func printer(c chan []int, b chan []int) {
     for {
          select {
          case msg := <-c:
               fmt.Println(msg)
               //time.Sleep(time.Second * 1)

          //Required to read the last b channel, so it will not stuck on b <- bucket
          case msg := <-b:
               fmt.Printf("last b %v\n", msg)
          }
     }
}

func main() {

     c := make(chan []int)
     bucket := make(chan []int)

     go runner(1, c, bucket)
     go runner(2, c, bucket)
     go runner(3, c, bucket)
     go runner(4, c, bucket)

     bucket <- []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
     ///close(bucket)

     go printer(c, bucket)

     var input string
     fmt.Scanln(&input)

}
