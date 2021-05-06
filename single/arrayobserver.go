package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type list struct {
	m  map[int]string
	mu sync.RWMutex
	br *bracket
}

func (l *list) add(idx int, v string) {
	l.m[idx] = v
}

type bracket struct {
	observee []int
}

func (b *bracket) observer(done <-chan bool, fnLookup func(int) string) {

	go func() {

		for {

			select {
			case <-done:
				return
			default:
				x := ""
				for _, k := range b.observee {
					x = x + fmt.Sprintf(" %d -> %s ", k, fnLookup(k))
				}
				fmt.Println(x)
			}

		}

	}()
}

func main() {

	//list := make(map[int]string)
	done := make(chan bool)
	l := &list{
		m: make(map[int]string),
	}
	l.br = &bracket{}
	l.br.observer(done, func(idx int) string {
		l.mu.Lock()
		defer l.mu.Unlock()
		return l.m[idx]
	})

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup, i int, l *list) {
			defer wg.Done()

			l.mu.Lock()
			l.add(i, randSeq(10))
			l.br.observee = append(l.br.observee, i)
			l.mu.Unlock()
			//	time.Sleep(time.Duration(rand.Intn(6-2)+2) * time.Second)
			fmt.Printf("Done %d\n", i)
		}(&wg, i, l)
	}

	wg.Wait()

	//fmt.Println("wait to done")
	time.Sleep(1 * time.Second)
	done <- true
	//	close(done)

	fmt.Println(l.br.observee)
	fmt.Println(l.m)

}
