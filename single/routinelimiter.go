package main

import (
	"fmt"
	"sync"
)

func main() {

	var apps []int
	for i := 1; i <= 450; i++ {
		apps = append(apps, i)
	}

	var wg sync.WaitGroup
	var counter int
	limiter := 200
	var m sync.Mutex
	var checker []int
	for _, app := range apps {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			//time.Sleep(time.Duration(i) * time.Second)
			m.Lock()
			checker = append(checker, i)
			m.Unlock()
			fmt.Printf("Print %d\n", i)
		}(&wg, app)

		counter++
		if limiter > 0 && counter == limiter {
			counter = 0
			wg.Wait()
			fmt.Println("*********************startnew")
		}
	}

	wg.Wait()

	fmt.Println("done")

	fmt.Println(len(checker))
}
