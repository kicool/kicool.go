package permutation 

import (
	"fmt"
	"testing"
)

func TestSimple(*testing.T) {
	a := []int{1, 2, 3}
	fmt.Printf("%v\n", a)
	for true == Next(a) {
		fmt.Printf("%v\n", a)
	}
}
