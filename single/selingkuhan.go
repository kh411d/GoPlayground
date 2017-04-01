package main

import (
	"fmt"
	"reflect"
)

type babe struct {
    baba anak
     	
}

func (s *babe) selingkuhan() {
 fmt.Println("iya ini selingkuhan")
 s.baba.bleh()
}

type anak struct {
	iya bool
	toSelingkuh interface{}
}

func (s *anak) bleh() {
 fmt.Println("iya ini selingkuhan selingkuhin anaknya")
}

func  (s *anak) boleh() {
    t := reflect.TypeOf(s)
    v := reflect.ValueOf(s)
    fmt.Printf("model: %+v %+v %+v\n", s, t, v)
}

func main() {
	fmt.Println("Hello, playground")
	
	babe := babe{}
	fmt.Println(babe)
	babe.selingkuhan()
	babe.baba.boleh()
}
