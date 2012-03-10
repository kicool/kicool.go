package main

import (
	"fmt"
	"math"
)

// what is the difference amongs Print* func()?
func main() {
	fmt.Printf("Now you have %g problems/n", math.Nextafter(2, 3))
	fmt.Println()
	fmt.Printf("Pi is ")
	fmt.Print(math.Pi)
	fmt.Print(".")
	fmt.Println()
}

