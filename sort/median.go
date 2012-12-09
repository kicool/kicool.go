package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	a     []int = []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	DEBUG bool
	kth   int = 1
)

func init() {
	flag.BoolVar(&DEBUG, "debug", false, "no debug by default")
	flag.IntVar(&kth, "kth", 1, "select kth element")
}

func swap(a []int, p1, p2 int) {
	t := a[p1]
	a[p1] = a[p2]
	a[p2] = t
}

/**
* In linear time, group the subarray ar[left, right] around a pivot
* element pivot=ar[pivotIndex] by storing pivot into its proper
* location, store, within the subarray (whose location is returned
* by this function) and ensuring that all ar[left,store) <= pivot and
* all ar[store+1,right] > pivot.
 */
func partition(a []int, l, r, pivotIndex int) int {
	if DEBUG {
		log.Println("<partition:", a, l, r, pivotIndex)
	}

	if l > r || pivotIndex < l || pivotIndex > r {
		panic("bad argument: out of index")
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
	return l + (r-l+1)/2
}

/**
* Average-case linear time recursive algorithm to find position of kth
* element in ar, which is modified as this computation proceeds.
* Note 1 <= k <= right-left+1. The comparison function, cmp, is
* needed to properly compare elements. Worst-case is quadratic, O(n^2).
 */
func selectKth(a []int, l, r, k int) int {
	if DEBUG {
		log.Println("selectKth:", l, r, k)
	}

	if l > r {
		panic("selectKth: out of index")
	}

	if k < 1 || k > (r-l+1) {
		panic("selectKth: out of index")
	}

	if l == r && k == 1 {
		return l
	}

	idx := selectPivotIndex(a, l, r)
	p := partition(a, l, r, idx)

	if p == (l + k - 1) {
		return p
	}

	if p < (l + k - 1) {
		return selectKth(a, p+1, r, k-(p-l+1))
	}

	return selectKth(a, l, p-1, k)

}

func medianSort(a []int, l, r int) {
	if DEBUG {
		log.Println("mediaSort:", l, r)
	}

	if r <= l {
		return
	}

	mid := (r - l + 1) / 2
	selectKth(a, l, r, mid)

	medianSort(a, l, l+mid-1)
	medianSort(a, l+mid+1, r)
}

func main() {
	flag.Parse()

	fmt.Println(a, len(a))

	medianSort(a, 0, len(a)-1)

	fmt.Println(a)
}
