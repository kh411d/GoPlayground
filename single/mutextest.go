package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	stateUnlocked uint32 = iota
	stateLocked
)

var (
	locker    = stateUnlocked
	errLocked = errors.New("lock woy")
)

type Setter struct {
	mu sync.Mutex
	v  string
}

func (c *Setter) Set(x string) {
	c.setWithAtomic(x)
	//c.setWithMutex(x)
}

func (c *Setter) setWithMutex(x string) {
	c.mu.Lock()
	fmt.Printf("set for %s\n", x)
	c.v = x
	time.Sleep(2 * time.Second)
	c.mu.Unlock()
}

func (c *Setter) setWithAtomic(x string) error {
	if !atomic.CompareAndSwapUint32(&locker, stateUnlocked, stateLocked) {
		return errLocked
	}

	defer atomic.StoreUint32(&locker, stateUnlocked)

	fmt.Printf("set for %s\n", x)
	c.v = x

	time.Sleep(2 * time.Second)

	return nil
}

func main() {
	c := &Setter{}

	go func() {
		fmt.Println("run satu")
		c.Set("satu")
		return
	}()

	go func() {
		fmt.Println("run dua")
		c.Set("dua")
	}()

	go func() {
		fmt.Println("run tiga")
		c.Set("tiga")
		return
	}()

	time.Sleep(7 * time.Second)

	//extra proces after 123 done
	go func() {
		fmt.Println("run empat")
		c.Set("empat")
		return
	}()
	time.Sleep(2 * time.Second)
}
