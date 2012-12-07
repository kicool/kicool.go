package main

import (
	"fmt"
)

var (
	a        []int = []int{1, 5, 6, 7, 2, -1, 0, 32, -9, 10, 5, -3}
	cmpcount       = 0
	swpcount       = 0
)

/*
http://en.wikipedia.org/wiki/Insertion_sort
*/
func insertion(a []int) {
	l := len(a)

	for i := 1; i < l; i++ {
		//make position i as a "hole"
		hole := i

		//save the hole value a[hole] to item
		item := a[hole]

		cmpcount++
		//keep moving the hole to next smaller index until A[hole - 1] is <= item
		for hole > 0 && a[hole-1] > item {
			// move hole to next smaller index
			a[hole] = a[hole-1]
			swpcount++
			hole--
		}

		//put item in the hole
		a[hole] = item
		swpcount++
	}

}

func main() {
	fmt.Println(a)

	insertion(a)

	fmt.Printf("cmp:%d swp:%d\n", cmpcount, swpcount)
	fmt.Println(a)
}
