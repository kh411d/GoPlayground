// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

func FindParents(k int, v *Child, lookup map[int]*Child) {
	if l, ok := lookup[k]; ok {
		v.Parents = append(v.Parents, masterType{
			ID:   l.ID,
			Name: l.Name,
		})
		FindParents(l.PID, v, lookup)
	}
	return
}

type Child struct {
	ID      int
	Name    string
	PID     int
	Parents []masterType
}

type masterType struct {
	ID   int
	Name string
}

func main() {
	data := map[int]*Child{
		1: {ID: 1, Name: "one", PID: 0},
		2: {ID: 2, Name: "two", PID: 1},
		3: {ID: 3, Name: "three", PID: 2},
		4: {ID: 4, Name: "four", PID: 1},
		5: {ID: 5, Name: "five", PID: 0},
	}

	for _, v := range data {
		FindParents(v.PID, v, data)
	}

	for _, v := range data {
		b, _ := json.Marshal(v)
		fmt.Printf("%s\n", string(b))
	}
}
