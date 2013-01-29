package main

import (
	"fmt"
	"log"
)

func bye(s string) {
	log.Println("bye:", s)
}

func foo(i int) {
	log.Println("IN Foo ", i)
	defer bye(fmt.Sprintf("Defer3 Loop%d", i))
	log.Println("OUT Foo ", i)

}

func main() {
	log.Println("Start")

	{
		log.Println("IN1")
		defer bye("Defer1")
		log.Println("OUT1")
	}

	for i := 0; i < 3; i++ {
		log.Println("IN2 ", i)
		defer bye(fmt.Sprintf("Defer2 Loop%d", i))
		log.Println("OUT2 ", i)
	}

	for i := 0; i < 3; i++ {
		foo(i)
	}

	log.Println("End")
}
