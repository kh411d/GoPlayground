package main

import (
	"fmt"
	"regexp"
)

func main() {
	//r := regexp.MustCompile(`(?i)Balance.*`)

	r := regexp.MustCompile(`(?i)(.*)-[0-9\.]+`)

	fmt.Printf("%q\n", r.FindSubmatch([]byte("sadfasdf-sdfsd-3.3.24")))
}
