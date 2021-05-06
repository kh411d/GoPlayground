package main

import (
	"fmt"
	"time"
)

type observee []int

func (o *observee) observer(fn func(observee), td time.Duration) *time.Ticker {
	done := make(chan struct{})
	t := time.NewTicker(td)
	//Start a go routine
	go func() {
		var pastLen int
		//Loop forever
		for {
			select {
			case <-done:
				return
			case <-t.C:
				nowLen := len(*o)
				if nowLen != pastLen {
					//Do something!
					fn(*o)
				}
				pastLen = nowLen
			}
		}
	}()
	return t
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

	t.Stop()
	fmt.Println("Ticker/Observer stopped")
	fmt.Println("Appending resumed, but observer will skip 4,5,6")
	for i := 4; i <= 6; i++ {
		theArray = append(theArray, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Restart/Reset the observer")
	fmt.Println("Appending resumed, observer will continue print the output")
	t.Reset(1 * time.Nanosecond)
	for i := 7; i <= 9; i++ {
		theArray = append(theArray, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("The final array: %v\n", theArray)
}

/**/
