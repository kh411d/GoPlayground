package main

import (
  "fmt"
  "time"
)

func inTimeSpan(start, end, check time.Time) bool {
  if check.Equal(start) || check.Equal(end) {
    fmt.Println("cekc")
    return true
  }
  return check.After(start) && check.Before(end)
}

func main() {

  layout := "2006-01-02 15:04:05"
  fmt.Println("Hello, playground")

  p := fmt.Println

  now := time.Now()
  //now,_ := time.Parse(layout, "2018-01-01 08:20:00")

  p(now.Zone())
  //_, tzOffset := now.Zone()

  p(now)

  tA, _ := time.ParseInLocation(layout, now.Format("2006-01-02")+" 08:00:00", now.Location())
  tB, _ := time.ParseInLocation(layout, now.Format("2006-01-02")+" 16:00:00", now.Location())
  p(tA)
  p(tB)
  if inTimeSpan(tA, tB, now) {
    p("wow")
  }

}
