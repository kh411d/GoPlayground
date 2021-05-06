package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var w sync.WaitGroup
	zcount := 0
	for i := 1; i <= 5; i++ {
		w.Add(1)
		go func(wg *sync.WaitGroup, cnt int) {
			defer wg.Done()
			x := make(chan bool)
			y := make(chan bool)
			z := make(chan bool)

			w.Add(1)
			go func(wg *sync.WaitGroup, x chan bool, y chan bool, z chan bool, cnt int) {
				defer wg.Done()

				select {
				case o := <-x:
					fmt.Printf("xxx %#v - %d\n", o, cnt)
					return
				case <-time.After(time.Duration(rand.Intn(6-2)+2) * time.Second):
					fmt.Printf("delay done - %d\n", cnt)
					y <- true
					close(y)
					for m := range z {
						zcount++
						fmt.Printf("zzz %#v - %d\n", m, cnt)
					}
				}

			}(wg, x, y, z, cnt)
			//time.Duration(rand.Intn(6-2)+2) * time.Second
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

			/* this will dead lock */
			// cancel := func(x chan bool, z chan bool) {
			// 	x <- true
			// 	close(x)
			// 	z <- true
			// 	close(z)
			// }

			// for n := range y {
			// 	fmt.Printf("yyy %#v\n", n)
			// 	cancel(x, z)
			// }

			/* this will not*/
			cancel := func(x chan bool) {
				x <- true
				close(x)
			}

			discard := func(z chan bool) {
				z <- true
				close(z)
			}

			go func(_x chan bool, _z chan bool) {
				cancel(_x)
			}(x, z)

			go func(y chan bool, z chan bool, cnt int) {
				for n := range y {
					fmt.Printf("yyy %#v - %d\n", n, cnt)
					discard(z)

					//_z <- true
					//close(_z)
				}
			}(y, z, cnt)

			// time.Sleep(3 * time.Second)
			// fmt.Printf("done - %d", cnt)

		}(&w, i)
	}
	w.Wait()

	fmt.Printf("z count %d\n", zcount)

}
