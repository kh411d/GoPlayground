// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x struct {
		val string `json:"ini_dia"`
	}
	j := `{"ini_dia": null}`
	err := json.Unmarshal([]byte(j), &x)
	fmt.Println(err)
	fmt.Printf("%#v", x)

}
