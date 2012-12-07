package main

import (
	"fmt"
)

var (
	a        []int = []int{1, 5, 6, 7, 2, -1, 0, 32, -9, 10, 5, -3}
	cmpcount       = 0
	swpcount       = 0
)

func swap(a []int, p1, p2 int) {
	t := a[p1]
	a[p1] = a[p2]
	a[p2] = t
}

/*
http://en.wikipedia.org/wiki/Selection_sort
*/
func selection(a []int) {
	l := len(a)

	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			cmpcount++
			if a[j] < a[min] {
				min = j
			}
		}

		if min != i {
			swap(a, i, min)
			swpcount++
		}
	}

}

func main() {
	fmt.Println(a)

	selection(a)

	fmt.Printf("cmp:%d swp:%d\n", cmpcount, swpcount)
	fmt.Println(a)
}
