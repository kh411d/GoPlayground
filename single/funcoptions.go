package main

import (
	"fmt"
)

type client struct {
	name string
	ini  hoho
}

func test(name string, opts ...func(c *client)) {
	c := client{}
	for _, opt := range opts {
		opt(&c)
	}
}

type hoho struct {
	name string
}

func main() {
	fmt.Println("Hello, playground")
	x := hoho{}
	x.name = "ini hoho"
	opt := func(c *client) {
		c.name = "jono"
		c.ini = x
		fmt.Println(c.ini)
	}
	test("joni", opt)
}
