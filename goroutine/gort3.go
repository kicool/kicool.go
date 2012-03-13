package main

import (
	"time"
	"fmt"
)

func IsReady(what string, minutes int64) {
	time.Sleep(minutes * 60 * 1e9) // Unit is nanosecs.
	fmt.Println(what, "is ready")
}

func main() {
	go IsReady("tea", 6)
	go IsReady("coffee", 2)
	fmt.Println("I'm waiting...")
}
