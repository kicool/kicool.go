package permutation

/* 
 * git://gist.github.com/305859.git
 */

import "sort"

func compare(a, b []int) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func nextPermutation(seq []int) bool {
	var j int = -1
	var m int = -1

	// Attempt to find the j index, this is the first index where it's value is
	// less than the after it (i.e. seq[j] < seq[j+1])
	for index := len(seq) - 2; index >= 0; index-- {
		if seq[index] < seq[index+1] {
			j = index
			break
		}
	}

	// If no index was found then finish
	if j == -1 {
		return false
	}

	// Find the m index, this is the first index where the value is greater than
	// the value at j
	for index := len(seq) - 1; index > j; index-- {
		if seq[index] > seq[j] {
			m = index
			break
		}
	}

	// Swap the values at the two indexes
	seq[j], seq[m] = seq[m], seq[j]

	// Order all of the values after j
	upper := seq[j+1:]
	sort.Ints(upper)

	for i, v := range upper {
		seq[j+i+1] = v
	}

	return true
}

func Next(seq []int) bool {
	current := make([]int, len(seq), len(seq))
	copy(current, seq)

	for i := 1; i < len(seq); i++ {
		success := nextPermutation(seq)
		if compare(current, seq) != 0 {
			return success
		}
	}

	return false
}
