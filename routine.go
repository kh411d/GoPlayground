package main 
import (
	"fmt"
)
func main() {
	// Create a channel to synchronize goroutines
     done := make(chan bool)
     x := false
     // Execute println in goroutine
     go func() {
          fmt.Println("goroutine message")
           for i:=1;i<10000;i = i+1 {}
    
          // Tell the main function everything is done.
          // This channel is visible inside this goroutine because
          // it is executed in the same address space.
          done <- true
     }()

     fmt.Println("main function message")
     fmt.Println(x)
     x = <-done // Wait for the goroutine to finish
     fmt.Println(x)
}