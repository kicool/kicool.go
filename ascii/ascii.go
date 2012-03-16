package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 0xff; i++ {
		fmt.Printf("0x%:", i)
	}
}
