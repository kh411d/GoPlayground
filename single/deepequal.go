package main

import (
  "fmt"
  "reflect"
)

type structA struct {
  Bs structB
  Cs []structC
}

type structB struct {
  ini string
}

type structC struct {
  itu string
}

func main() {
  z := new(structA)
  z.Bs = structB{ini: "ini b"}
  z.Cs = append(z.Cs, structC{itu: "ini c"})

  y := new(structA)
  y.Bs = structB{ini: "ini b"}
  y.Cs = append(y.Cs, structC{itu: "ini d"})

  x := reflect.DeepEqual(z, y)
  fmt.Println(x)
}
