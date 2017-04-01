package main

import (
	"fmt"
	"reflect"
)

/*type builderInterface interface {
	setCache(s string) string
}

type builder struct {
 factory builderInterface
}

func (self *builder) get(s string) builderInterface {
	if(s == "binatang"){
		self.factory = new(binatang)
	 	return self.factory
	}

	return self.factory 
}*/

type all struct {
	anoa binatang
}

func getAll(s string) *all{
	c := new(all)
	//fmt.Println(reflect.ValueOf(c))

for i := 0; i < reflect.Indirect(reflect.ValueOf(c)).NumField(); i++ {
	  o := reflect.Indirect(reflect.ValueOf(c)).Field(i)
	  
		fmt.Println("Field:", o)
	}
// List methods of s.context
	for i := 0; i < reflect.TypeOf(c).NumMethod(); i++ {
		fmt.Println("Method:", reflect.TypeOf(c).Method(i).Name)
	}	


	return c
	/*if method := reflect.ValueOf(c).MethodByName(s); !method.IsNil() {
		return method.Call()
	} else {
		return c
	}*/
}

type binatang struct {}

func (self *binatang) Gunung() string {
  return "ini gunung"
}

func (self *binatang) setCache(s string) string{
	return "ini dicache"
} 

func main() {
	fmt.Println("Hello, playground")
	
	//x := builder{&binatang{}}
	//x := builder{}
	//c := new(builder).get("binatang").Gunung()
	c := getAll("anoa")

	fmt.Println(c.anoa.Gunung())
}
