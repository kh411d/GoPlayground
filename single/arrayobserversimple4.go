package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type observee []string

func (o *observee) observer(fn func(int, string, *sync.Mutex), td time.Duration) (*sync.Mutex, chan<- struct{}) {
	var mu sync.Mutex
	done := make(chan struct{})
	//Start a go routine
	go func(mu *sync.Mutex) {
		//Forever loop through each array element
		nextIdx := 0
		for {
			select {
			case <-done:
				return
			case <-time.NewTicker(td).C:
				if x := *o; x != nil {
					fn(nextIdx, x[nextIdx], mu)
					nextIdx++
					if nextIdx == len(x) {
						nextIdx = 0
					}
				}
			}
		}
	}(&mu)

	return &mu, done
}

type work struct {
	wg *sync.WaitGroup
}

func (w *work) countdown(workerID int) string {
	// return "[worker id]-[time(sec) to launch]-[timestamp added]"
	return fmt.Sprintf("%d-%d-%d", workerID, rand.Intn(5), time.Now().UnixNano())
}
func (w *work) ready(v string) (launchID int, ok bool) {
	id, tx, ts := w.data(v)
	if id > 0 && time.Now().Sub(time.Unix(0, ts)).Seconds() > float64(tx) {
		launchID = id
		ok = true
	}
	return
}
func (w *work) launch(i int) {
	w.wg.Add(1)
	go func(i int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("worker %d is launched\n", i)
	}(i, w.wg)
}
func (w *work) data(v string) (id int, tx int, ts int64) {
	if v == "" {
		return
	}
	s := strings.Split(v, "-")
	id, _ = strconv.Atoi(string(s[0]))
	tx, _ = strconv.Atoi(string(s[1]))
	ts, _ = strconv.ParseInt(string(s[2]), 10, 64)
	return
}

func main() {
	var theArray observee
	var wg sync.WaitGroup
	worker := &work{
		wg: &wg,
	}

	mu, t := theArray.observer(func(i int, v string, mu *sync.Mutex) {
		mu.Lock()
		if id, ok := worker.ready(v); ok {
			theArray[i] = ""
			worker.launch(id)
		}
		mu.Unlock()
	}, 1*time.Nanosecond) //observing each element per 1 Nanosecond

	for i := 1; i <= 10; i++ {
		go func(mu *sync.Mutex, i int) {
			mu.Lock()
			theArray = append(theArray, worker.countdown(i))
			mu.Unlock()
		}(mu, i)
	}

	wg.Wait()

	time.Sleep(5 * time.Second)
	t <- struct{}{}

}

/**
type observee []int

func (o *observee) observer(fn func(observee), td time.Duration) chan struct{} {
    done := make(chan struct{})
    go func() {
        for {
            select {
            case <-done:
                return
            case <-time.NewTicker(td).C:
                fn(*o)
            }
        }
    }()
    return done
}

func main() {
    var theArray observee

    t := theArray.observer(func(o observee) {
        //Current observed value
        fmt.Printf("Observed %v\n", o)
    }, 1*time.Nanosecond)

    pause := func() {
        time.Sleep(1 * time.Second)
    }

    //Append
    theArray = append(theArray, 1)
    pause()

    theArray = append(theArray, 2)
    pause()

    theArray = append(theArray, 3)
    pause()

    //Update by index
    theArray[2] = 4
    pause()

    theArray[1] = 4
    pause()

    theArray[0] = 4
    pause()

    t <- struct{}{}

    fmt.Println(theArray)

}
/**/
