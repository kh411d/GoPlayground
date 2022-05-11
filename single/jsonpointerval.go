package main

import (
	"encoding/json"
	"fmt"
)

type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

type emptyval struct {
	val
	Price *float64 `json:"price"`
}

type val struct {
	Price float64 `json:"price"`
}

func main() {
	data := `{"price": 32}`
	var x emptyval
	json.Unmarshal([]byte(data), &x)

	fmt.Printf("%#v", *x.Price)

	var s NullFloat64
	fmt.Printf("%#v", s)

	var z float64
	z = 234
	b, _ := json.Marshal(emptyval{
		Price: &z,
	})
	fmt.Printf("\n%#v", string(b))

}
