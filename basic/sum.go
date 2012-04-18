package main

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}

	return s
}

func main() {
	a0 := [4]int{5, 7, 8, 0}
	a1 := [...]int{1, 2, 3, 4}

	s0 := sum(a0[:])
	s1 := sum(a1[:])
	s2 := sum([]int{1, 2, 3})

	// Build error
	// invalid operation [3]int literal[:] (slice of unaddressable value)
	// s3 := sum([...]int{1,2,3}[:])

	println(s0, s1, s2)
}
