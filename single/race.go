package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	xm["test"].inc()
	wg.Done()
}

var xm = map[string]*corp{
	"test": &corp{},
}

type corp struct{}

func (*corp) inc() {
	x = x + 1
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
