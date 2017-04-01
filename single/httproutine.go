package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	redis "gopkg.in/redis.v5"
)

// key cZlOSQyv0cTFMnDQQy3vZZswotk8BtcmQFlZVEzk

var lock = sync.RWMutex{}

func createRedisConnection() *redis.Client {
	lock.Lock()
	defer lock.Unlock()

	conn := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		MaxRetries:  2,
		IdleTimeout: 5 * time.Minute,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		panic(err)
	}

	return conn
}

func request(sol int) { //[]byte {
	redisConn := createRedisConnection()
	url := fmt.Sprintf("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=%s&api_key=cZlOSQyv0cTFMnDQQy3vZZswotk8BtcmQFlZVEzk", strconv.Itoa(sol))
	fmt.Println(url)
	var r *http.Request
	r, _ = http.NewRequest("GET", url, nil)
	r.Header.Add("Content-Type", "application/json")

	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		//return nil, errors.New("Service is unreachable. please try again later.")
		fmt.Println(err)
	}

	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)

	if err = redisConn.Set(strconv.Itoa(sol), string(response), 120*time.Minute).Err(); err != nil {
		log.Printf("redis set error: %+v\n", err.Error())
	}
	//fmt.Println(err)
	//return response //, err
}

func getChannel(i int) { //<-chan []byte {
	//c := make(chan []byte, 20)
	go func() {
		//c <- request(i)
		request(i)
	}()
	//return c
}

func main() {

	//r, err := request(100)
	//fmt.Println(string(r))
	//fmt.Println(err)

	/*done := make(chan bool)
	go func() {
		for i := 100; i <= 120; i++ {
			request(i)

		}
		done <- true
	}()
	x := <-done // Wait for the goroutine to finish
	fmt.Println(x)*/

	/*for i := 100; i <= 120; i++ {
		go func(i int) {
			request(i)
		}(i)
	}

	time.Sleep(time.Second * 5)*/

	/*go func() {
		for i := 100; i <= 120; i++ {

			request(i)

		}
	}()

	time.Sleep(time.Second * 60)*/

	done := make(chan int)

	for i := 100; i <= 120; i++ {
		go func(i int, c chan<- int) {
			request(i)
			c <- i
		}(i, done)
	}

	/*for {
		select i <-done
		fmt.Println(<-done)
	}*/
	//var x chan int
	for {
		select {
		case <-done:
			fmt.Println(<-done)
		}
	}

}
