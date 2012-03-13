package main

import (
	"time"
)

func server(i int) {
	for {
		print(i)
		time.Sleep(1*1e9)
	}
}

/*
* this will no output if default: export GOMAXPROCS=1
 */
func f1() {
	go server(1)
	go server(2)
	for {} 
}

func f3() {
	go server(1)
	go server(2)
	for {
		time.Sleep(1*1e6)
	} 
}

func f2() {
	go server(1)
	go server(2)
	time.Sleep(6*1000*1000) 
}

func main() {
	f3()
	//f2()
}
