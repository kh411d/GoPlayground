package main

import (
     "fmt"
     "time"
)

func runner(c chan []int, bucket *ember) {
     var temp []int

     for k, v := range bucket.isi {
          if v != 0 {
               temp = append(temp, v)
               bucket.isi[k] = 0
          }
          if len(temp) == 5 {
               break
          }
     }
     fmt.Println(bucket)

     c <- temp

}

func printer(c chan []int, bucket *ember) {
     for {
          msg := <-c
          // fmt.Println(msg)
          fmt.Println(msg)

          fmt.Println("")
          time.Sleep(time.Second * 1)

     }
}

type ember struct {
     isi []int
}

func main() {

     isinya := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

     bucket := &ember{isi: isinya}

     c := make(chan []int, 5)

     go runner(c, bucket)
     go runner(c, bucket)
     go runner(c, bucket)
     go runner(c, bucket)

     go printer(c, bucket)

     var input string
     fmt.Scanln(&input)

}
