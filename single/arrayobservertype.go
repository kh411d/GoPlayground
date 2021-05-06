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

type observee []int

func (o *observee) observer(done <-chan bool, td time.Duration, fnLookup func(observee)) {
	ticker := time.NewTicker(td)
	go func() {
		var pastLen int
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				nowLen := len(*o)
				if pastLen != nowLen {
					fnLookup(*o)
				}

				pastLen = nowLen
			}
		}
	}()
}

func (o observee) all() []int {
	return o
}

func main() {

	done := make(chan bool)
	var br observee
	//br := make(observee, 10)
	br.observer(done, 1*time.Nanosecond, func(v observee) {
		fmt.Printf("replay %d\n", v)
	})
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup, i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			//time.Sleep(time.Duration(rand.Intn(6-2)+2) * time.Second)
			br = append(br, i)
			fmt.Printf("Done %d\n", i)
		}(&wg, i)
	}

	wg.Wait()
	time.Sleep(1 * time.Second)
	done <- true
	//close(done)

	fmt.Println(br)

}
