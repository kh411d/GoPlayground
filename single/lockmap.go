package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

//Shelf temporary order place before delivery
type Shelf struct {
	mu      sync.RWMutex
	counter map[string]int
	items   map[string]string
	ids     []string

	hotshelf *hotshelf
}

type hotshelf struct {
	ids []string
}

func (h hotshelf) add(v string) {
	h.ids = append(h.ids, v)
}
func (h hotshelf) delete(v string) {
	i := sort.SearchStrings(h.ids, v)
	h.ids[i] = h.ids[len(h.ids)-1] // Copy last element to index i.
	h.ids[len(h.ids)-1] = ""       // Erase last element (write zero value).
	h.ids = h.ids[:len(h.ids)-1]
}

func (s *shelf) append(x string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[x] = x
	s.hotshelf.add(x)
	s.counter[x] = s.counter[x] + 1
}

func (s *shelf) delete(x string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, x)
	s.counter[x] = s.counter[x] - 1
	s.hotshelf.add(x)
}

func (s *Shelf) set(v int) {
	//s.mu.Lock()
	//defer s.mu.Unlock()
	fmt.Printf("%#v\n", s)
	//s.counter["a"] = v
	s.sets(v)

	time.Sleep(5 * time.Second)
}

func (s *Shelf) sets(v int) {

	s.counter["a"] = v

}

func (s *Shelf) add(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = append(s.ids, v)

	time.Sleep(5 * time.Second)
}

func (s *Shelf) get() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := s.counter

	time.Sleep(5 * time.Second)
	return v
}

func main() {
	s := &Shelf{
		items: make(map[string]string),
		counter: map[string]int{
			"a": 10,
			"b": 10,
		},
		ids: []string{"a", "b", "c"},
	}

	s.set(8)

	go func(s *Shelf) {
		//time.Sleep(1 * time.Second)
		fmt.Println("try set a")
		s.set(5)
	}(s)

	//s.set(6)

	time.Sleep(3 * time.Second)
	fmt.Printf("%#v\n", s)

}
