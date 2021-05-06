package main

import (
	"fmt"
	"errors"
)

type customErr struct {
   messg string
	data interface{}
	error
}

func(c *customErr) Error() string {
	return c.messg
}


func main() {
	fmt.Println("Hello, playground")
	
	x := errors.New("default error")
	
	fmt.Println(customErr(x))
}

