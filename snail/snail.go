// http://hackgolang.blogspot.com/2010/09/snail-in-golang.html
package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// Determine the matrix size, n.
	var n int = 1
	flag.Parse()
	if len(flag.Args()) > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}
	if n < 1 {
		n = 1
	}

	// No need to be complicated for n=1
	if n == 1 {
		fmt.Println(n)
		return
	}

	// Create the matrix
	values := make([][]int, n)
	for i := 0; i < n; i++ {
		values[i] = make([]int, n)
	}

	// Fill the matrix with values
	avalue := 0
	for ifill := 0; ifill < n; ifill++ {
		fstart := ifill
		fend := n - ifill

		// upper x forward sweep
		for ixf := fstart; ixf < fend-1; ixf++ {
			avalue++
			values[fstart][ixf] = avalue
		}

		// right y forward sweep
		for iyf := fstart; iyf < fend-1; iyf++ {
			avalue++
			values[iyf][fend-1] = avalue
		}

		// lower x backward sweep
		for ixb := fend - 1; ixb > fstart; ixb-- {
			avalue++
			values[fend-1][ixb] = avalue
		}

		// left y backward sweep
		for iyb := fend - 1; iyb > fstart; iyb-- {
			avalue++
			values[iyb][fstart] = avalue
		}
	}

	// Take care of the last value for odd n
	if avalue < n*n {
		avalue++
		values[n/2][n/2] = avalue
	}

	// Ensure the same spacing for each element
	places := int(n * n / 10)
	entry_width := 1
	for places > 0 {
		entry_width++
		places /= 10
	}
	format := fmt.Sprintf("%%%dd ", entry_width)

	// Print out the snail
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(format, values[i][j])
		}
		fmt.Println("")
	}

}
