package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

var (
	a     []int = []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	DEBUG bool
)

func init() {
	flag.BoolVar(&DEBUG, "debug", false, "no debug by default")
}

func swap(a []int, p1, p2 int) {
	t := a[p1]
	a[p1] = a[p2]
	a[p2] = t
}

func partition(a []int, l, r, pivotIndex int) int {
	if l > r || pivotIndex < l || pivotIndex > r {
		panic("bad argument: out of index")
	}

	if l == r {
		return l
	}

	if DEBUG {
		log.Println("<partition:", a, l, r, pivotIndex)
	}

	pivot := a[pivotIndex]

	//move pivot to the end of the array
	swap(a, pivotIndex, r)

	/* all values <= pivot are moved to front of array and pivot inserted just after them. */
	store := l
	for i := l; i < r; i++ {
		if a[i] < pivot {
			swap(a, i, store)
			store++
		}
	}

	swap(a, r, store)

	if DEBUG {
		log.Println(">partition:", a, store)
	}

	return store
}

/* select [l, r] */
func selectPivotIndex(a []int, l, r int) int {
	if l >= r {
		return l
	}

	return l + rand.Intn(r-l)
}

func quickSort(a []int, l, r int) {
	if DEBUG {
		log.Println("quickSort:[D1]", a, l, r)
	}

	if l >= r {
		return
	}

	pivotIndex := selectPivotIndex(a, l, r)
	if DEBUG {
		log.Println("quickSort:[D2]", pivotIndex)
	}

	pivotIndex = partition(a, l, r, pivotIndex)
	if DEBUG {
		log.Println("quickSort:[D3]", pivotIndex)
	}

	if l < pivotIndex-1 {
		quickSort(a, l, pivotIndex-1)
	}
	if pivotIndex+1 < r {
		quickSort(a, pivotIndex+1, r)
	}
}

func main() {
	flag.Parse()

	fmt.Println(a, len(a))

	quickSort(a, 0, len(a)-1)

	fmt.Println(a)
}
