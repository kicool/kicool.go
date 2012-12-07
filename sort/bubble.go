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
http://en.wikipedia.org/wiki/Bubble_sort
*/
func bubble(a []int) {
	l := len(a)

	for {
		swapped := false

		for j := 1; j < l; j++ {
			cmpcount++
			if a[j-1] > a[j] {
				swap(a, j-1, j)
				swapped = true
				swpcount++
			}
		}

		if !swapped {
			break
		}
	}

	return
}

func bubble2(a []int) {
	l := len(a)

	for {
		swapped := false

		for j := 1; j < l; j++ {
			cmpcount++
			if a[j-1] > a[j] {
				swap(a, j-1, j)
				swapped = true
				swpcount++
			}
		}

		if !swapped {
			break
		}

		l--
	}

	return
}

func bubble3(a []int) {
	l := len(a)

	for {
		newn := 0

		for j := 1; j < l; j++ {
			cmpcount++
			if a[j-1] > a[j] {
				swap(a, j-1, j)
				newn = j
				swpcount++
			}
		}

		if newn == 0 {
			break
		}

		l = newn
	}

	return
}

func main() {
	fmt.Println(a)

	bubble3(a)

	fmt.Printf("cmp:%d swp:%d\n", cmpcount, swpcount)
	fmt.Println(a)
}
