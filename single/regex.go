package main

import (
	"fmt"
	"regexp"
)

func main() {
	//r := regexp.MustCompile(`(?i)Balance.*`)

	fmt.Println(regexp.MustCompile(`(?i)Balance.*`).MatchString("balance2342_s4f3433"))
}
