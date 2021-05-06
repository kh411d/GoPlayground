package main

import (
	"fmt"
	"time"
)

type observee []int

func (o *observee) observer(fn func(observee), td time.Duration) chan struct{} {
	done := make(chan struct{})
	//Start a go routine
	go func() {
		var pastLen int
		//Loop forever
		for {
			select {
			case <-done:
				return
			case <-time.NewTicker(td).C:
				nowLen := len(*o)
				if nowLen != pastLen {
					//Do something!
					fn(*o)
				}
				pastLen = nowLen
			}
		}
	}()
	return done
}

/**
func main() {
	var theArray observee
	var wg sync.WaitGroup
	var mu sync.Mutex

	t := theArray.observer(func(o observee) {
		fmt.Printf("Observe %v\n", o)
	}, 1*time.Nanosecond)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup, i int) {
			defer wg.Done()
			//make random pause for each go routine before appending
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			//append theArray with lock
			mu.Lock()
			theArray = append(theArray, i)
			mu.Unlock()
		}(&wg, i)
	}

	wg.Wait()
	//kill the observer
	t <- struct{}{}
	fmt.Printf("The final array: %v\n", theArray)
}

/**/

/**/
func main() {
	var theArray observee

	t := theArray.observer(func(o observee) {
		fmt.Printf("Observe %v\n", o)
	}, 1*time.Nanosecond)

	for i := 1; i <= 3; i++ {
		theArray = append(theArray, i)
		time.Sleep(1 * time.Second)
	}

	t <- struct{}{}
	fmt.Println("Ticker/Observer stopped")

	fmt.Println("resume appending, observer will not print any updates")
	for i := 4; i <= 6; i++ {
		theArray = append(theArray, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("resume appending, observer will not print any updates")
	for i := 7; i <= 9; i++ {
		theArray = append(theArray, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("The final array: %v\n", theArray)
}

/**/
