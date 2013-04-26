package main

// #include <stdlib.h>
import "C"


func main() {
	i := C.random()
	print(i)
}
