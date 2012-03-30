package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 0xff; i++ {
		fmt.Printf("0x%0X %c  ", i, i)
		if 0 == i % 8 {
			fmt.Println()
		}
	}
}
