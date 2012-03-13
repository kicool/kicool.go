package main

/*
 * demo a way to write a daemon
 */
import (
	"os"
)

// Receives the change in the number of goroutines
var goroutineDelta = make(chan int)
var needToCreateANewGoroutine = true

func main() {
	go forever()

	numGoroutines := 0
	for diff := range goroutineDelta {
		numGoroutines += diff
		if numGoroutines == 0 {
			os.Exit(0)
		}
	}
}

// Conceptual code
func forever() {
	for {
		if needToCreateANewGoroutine {
			// Make sure to do this before "go f()", not within f()
			goroutineDelta <- +1

			go f()
		}
	}
}

func f() {
	// When the termination condition for this goroutine is detected, do:
	goroutineDelta <- -1
}
