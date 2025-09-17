package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
)

func Process(rs *redsync.Redsync, log string) {
	// Create a mutex with WithTries(1) to disable retries
	mutex := rs.NewMutex("my-resource-lock" /**/, redsync.WithTries(1) /**/)

	// Attempt to acquire the lock without retries
	if err := mutex.LockContext(context.Background()); err != nil {
		if err == redsync.ErrFailed {
			fmt.Println(log + " Failed to acquire lock immediately (lock already held).")
		} else {
			fmt.Printf("%s Error acquiring lock: %v\n", log, err)
		}
		return
	}

	fmt.Println(log + " - Lock acquired")

	// Simulate work
	time.Sleep(3 * time.Second)

	// Release the lock
	if _, err := mutex.Unlock(); err != nil {
		fmt.Printf("%s - Error unlocking: %v\n", log, err)
	} else {
		fmt.Println(log + " - Lock released.")
	}
}

func main() {
	// Create a Redigo connection pool
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}

	// Create a Redsync instance
	rs := redsync.New(redigo.NewPool(pool))

	go Process(rs, "proses satu")
	go Process(rs, "proses dua")
	go Process(rs, "proses tiga")

	time.Sleep(15 * time.Second)
}
