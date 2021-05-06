package arrobs

import (
	"fmt"
	"testing"
	"time"
)

var result []int

func BenchmarkArrobs(b *testing.B) {
	var r []int
	for n := 0; n < b.N; n++ {
		var theArray observee

		t := theArray.observer(func(o observee) {
			fmt.Printf("Observe %v\n", o)
		}, 1*time.Nanosecond)

		for i := 1; i <= 50; i++ {
			theArray = append(theArray, i)
			//time.Sleep(1 * time.Second)
		}

		t <- struct{}{}
		r = theArray
	}

	result = r
	fmt.Printf("Finish %v\n", result)
}
