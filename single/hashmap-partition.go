package main

import (
	"fmt"
	"sync"
)

type Partition struct {
	sync.RWMutex
	m map[string]string
}

const partCount = 64

var m [partCount]Partition

// func Find(k string) string {
// 	idx := hash(k) % partCount
// 	part := &m[idx]
// 	part.RLock()
// 	v := part.m[k]
// 	part.RUnlock()
// 	return v
//}

func main() {
	m = Partition{
		m: map[string]string{
			"one": "isone",
			"two": "istwo",
		},
	}

	//x := Find("two")
	fmt.Println(len(m))
}
