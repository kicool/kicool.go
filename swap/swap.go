package main

import "fmt"

func swap1(x, y int) (int, int) {
	return y, x
}

// named result var
func swap2(x, y int) (_x, _y int) {
	_x, _y = y, x
	return
}

func main() {
	x, y := 1, 2
	fmt.Println("before swap: ", x, y)
	x, y = swap1(x, y)
	fmt.Println("after  swap1: ", x, y)
	x, y = swap2(x, y)
	fmt.Println("after  swap2: ", x, y)

}
