package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"
)

type request struct {
	name  string
	delay time.Duration
}

var Queue = make(chan request)

var Pool = make(chan chan request)

var TotalRequest int = 5
var TotalWorker int = 3

func main() {

	go func() {
		for i := 1; i <= TotalRequest; i++ {
			Queue <- request{
				name:  "request " + strconv.Itoa(i),
				delay: time.Second * 5,
			}
		}
	}()

	for j := 1; j <= TotalWorker; j++ {
		go func(workerNo int) {
			for {
				fmt.Println("spawn worker " + strconv.Itoa(workerNo))

				worker := make(chan request)
				Pool <- worker
				select {
				case o := <-worker:
					time.Sleep(o.delay)
					fmt.Println("Finish ", o.name)
					fmt.Println()
				}
				fmt.Println("Pool NumGoroutine ", runtime.NumGoroutine())
			}
		}(j)
	}

	go func() {
		for {
			select {
			case req := <-Queue:
				go func() {
					worker := <-Pool
					fmt.Println("Start " + req.name)
					worker <- req
					fmt.Println("Worker assign NumGoroutine ", runtime.NumGoroutine())
				}()
			}
		}
		fmt.Println("Worker all NumGoroutine ", runtime.NumGoroutine())
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		Queue <- request{
			name:  r.FormValue("name"),
			delay: time.Second * 5,
		}
	})

	/*http.HandleFunc("/debug/pprof/", func(w http.ResponseWriter, r *http.Request) {
		pprof.Index(w, r)
	})*/

	/*http.HandleFunc("/debug/pprof/profile", func(w http.ResponseWriter, r *http.Request) {
		pprof.Profile(w, r)
	})
	*/
	/*http.HandleFunc("/debug/pprof/:name", func(w http.ResponseWriter, r *http.Request) {
		pprof.Handler("heap").ServeHTTP(w, r)
	})
	*/
	if err := http.ListenAndServe("localhost:3003", http.DefaultServeMux); err != nil {
		fmt.Println(err.Error())
	}
}
