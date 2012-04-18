package scanner

import (
	"testing"
	"strings"
	"fmt"
)

// Tests int64 and uint64
func TestScanner1(t *testing.T) {
	sc := NewScanner(strings.NewReader("1  -662   4000000000 4000000000000000000  "))

	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if sc.NextInt64() != int64(1) {
		t.Errorf("wrong value read")
	}
	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if sc.NextInt64() != int64(-662) {
		t.Errorf("wrong value read")
	}
	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if sc.NextInt64() != int64(4000000000) {
		t.Errorf("wrong value read")
	}
	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if !sc.HasNextInt64() {
		t.Errorf("should have next int64")
	}
	if sc.NextInt64() != 4000000000000000000 {
		t.Errorf("wrong value read")
	}
	if sc.HasNextInt64() {
		t.Errorf("should NOT have next int64")
	}

	sc = NewScanner(strings.NewReader(" 1 8000000000000000000 -6 9999999999999999999"))

	if !sc.HasNextUint64() {
		t.Errorf("should have next uint64")
	}
	if sc.NextUint64() != 1 {
		t.Errorf("wrong value read")
	}
	if !sc.HasNextUint64() {
		t.Errorf("should have next uint64")
	}
	if sc.NextUint64() != 8000000000000000000 {
		t.Errorf("wrong value read")
	}
	if sc.HasNextUint64() {
		t.Errorf("should NOT have next uint64")
	}
	if sc.HasNextUint64() {
		t.Errorf("should NOT have next uint64")
	}
	if sc.HasNextUint64() {
		t.Errorf("should NOT have next uint64")
	}

	if sc.NextInt64() != -6 {
		t.Errorf("wrong value read")
	}
	if sc.HasNextInt64() {
		t.Errorf("should NOT have next int64")
	}
}

func TestScanner2(t *testing.T) {
	input := `    125  6   00 9139081 -1309714037 1037104 -0183

	    ` + "     \t\t\t   " + `
	   091 apa        kabar 0      `

	{
		s := strings.NewReader(input)
		sc := NewScanner(s)

		c := 0
		for sc.HasNextInt() {
			fmt.Printf("int:#%d#\n", sc.NextInt())
			c++
		}

		if c != 8 {
			t.Errorf("should have 8 ints")
		}
	}
	{
		sc := NewScannerString(input)

		fmt.Printf("uint:#%d#\n", sc.NextUint())
		fmt.Printf("uint:#%d#\n", sc.NextUint())
		fmt.Printf("uint:#%d#\n", sc.NextUint())
		fmt.Printf("uint:#%d#\n", sc.NextUint())

		c := 0
		for sc.HasNextLine() {
			fmt.Println("line:#" + sc.NextLine() + "#")
			c++
		}

		if c != 4 {
			t.Errorf("should have 4 lines")
		}
	}
	{
		sc := NewScannerString(input)

		c := 0
		for sc.HasNext() {
			x := sc.Next()
			fmt.Println("token:#" + x + "#")
			if x != strings.TrimSpace(x) {
				t.Errorf("should not have leading/trailing spaces")
			}
			c++
		}

		if c != 11 {
			t.Errorf("should have 11 tokens")
		}
	}
}

func testHasNext(t *testing.T, s string, _int, _int64, _uint, _uint64, _line, _token bool) {
	sc := NewScannerString(s)

	if sc.HasNextInt() != _int {
		t.Errorf("%s has int should be %v", s, _int)
	}
	if sc.HasNextInt64() != _int64 {
		t.Errorf("%s has int64 should be %v", s, _int64)
	}
	if sc.HasNextUint() != _uint {
		t.Errorf("%s has uint should be %v", s, _uint)
	}
	if sc.HasNextUint64() != _uint64 {
		t.Errorf("%s has uint64 should be %v", s, _uint64)
	}
	if sc.HasNextLine() != _line {
		t.Errorf("%s has line should be %v", s, _line)
	}
	if sc.HasNext() != _token {
		t.Errorf("%s has token should be %v", s, _token)
	}
}

func TestScanner3(t *testing.T) {
	testHasNext(t, "123", true, true, true, true, true, true)
	testHasNext(t, "-123", true, true, false, false, true, true)
	testHasNext(t, "123456789123456789", false, true, false, true, true, true)
	testHasNext(t, "-123456789123456789", false, true, false, false, true, true)
	testHasNext(t, " ", false, false, false, false, true, false)
	testHasNext(t, "", false, false, false, false, false, false)
}

func TestScanner4(t *testing.T) {
	right := []interface{}{1, 5, -2147483648, 2147483647, uint(4294967295), uint(0), "hore"}

	s := ""
	for _, v := range right {
		s += fmt.Sprint(v) + " "
	}

	sc := NewScannerString(s)
	i := 0

	check := func(got interface{}) {
		if got != right[i] {
			t.Errorf("got (%T)%v, should be (%T)%v", got, got, right[i], right[i])
		}
		i++
	}

	for sc.HasNextInt() {
		check(sc.NextInt())
	}
	for sc.HasNextUint() {
		check(sc.NextUint())
	}
	for sc.HasNext() {
		check(sc.Next())
	}
}
