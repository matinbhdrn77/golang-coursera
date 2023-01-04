// main() executes two anonymous functions. Keyword 'go' makes them run
// concurrently, in goroutines.
// Go runtime scheduler schedules goroutines in a non-deterministic way so
// there are multiple possible interleavings - each time we run the application
// scheduler might change the order of their execution.
// Sometimes function which increments x will be executed first while sometimes
// the function which prints out its value.
// Because both functions in goroutines are communicating through variable x
// there is a chance of race condition - the order of accessing variable x
// is non-deterministic.

package main

// Compulsory package,
// the only one generating an executable

import (
	"fmt"
	"time"
)

var x int

// Script containing two goroutines.
// One goroutine implements an increment of the variable.
// The other one prints the value of the variable on screen.
func main() {
	for i := 0; i < 100; i++ {
		go func() {
			x++
		}()

		go func() {
			fmt.Println("Value of x: ", x)
		}()
	}

	// prevent program termination before goroutines return
	time.Sleep(1 * time.Second)
}
