// canner is a simple parser for primitive types like int, float, and strings.
// It takes an io.Reader as input and produces primitive variables.
package scanner

import (
	"bufio"
	"bytes"
	"container/list"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// Scanner is an object that reads a byte-oriented stream and converts human-readable values
// to integers, tokens, and lines. Create a Scanner using NewScanner or NewScannerString.
type Scanner struct {
	in   *bufio.Reader
	list *list.List
}

type sd struct {
	s     string
	delim int
	err   error
}

// Creates a new Scanner with io.Reader as input source
func NewScanner(r io.Reader) *Scanner {
	sc := Scanner{in: bufio.NewReader(r)}
	sc.list = list.New()
	return &sc
}

// Creates a new Scanner with a string as input source
func NewScannerString(s string) *Scanner {
	sc := Scanner{in: bufio.NewReader(strings.NewReader(s))}
	sc.list = list.New()
	return &sc
}

func (this *Scanner) nextToken() (s string, delim int, err error) {
	buf := bytes.NewBufferString("")

	for {
		if c, _, e := this.in.ReadRune(); e == nil {
			if unicode.IsSpace(c) {
				s = buf.String()
				delim = int(c)
				return // (token, delim, nil)
			} else {
				buf.WriteString(string(c))
			}
		} else {
			if e == io.EOF {
				if buf.Len() > 0 {
					s = buf.String()
					return // (token, 0, nil)
				}
			}
			err = e
			return // ("", 0, os.EOF)
		}
	}

	return
}

func (this *Scanner) nextBuffedToken(after *list.Element) *list.Element {
	if after == nil { // get the first one only
		if this.list.Len() == 0 {
			s, delim, err := this.nextToken()
			next := sd{s: s, delim: delim, err: err}
			this.list.PushBack(next)
		}
		return this.list.Front()
	}

	// get the one after /after
	if after.Next() == nil {
		s, delim, err := this.nextToken()
		next := sd{s: s, delim: delim, err: err}
		this.list.PushBack(next)
	}
	// debug assert
	if after.Next() == nil {
		panic("after.Next() == nil")
	}
	return after.Next()
}

func (this *Scanner) popBuff() {
	if this.list.Len() == 0 {
		panic("should not pop list with len 0")
	}
	this.list.Remove(this.list.Front())
}

func (this *Scanner) nextSth(converter func(string) (int64, error)) int64 {
	for {
		next := this.nextBuffedToken(nil).Value.(sd)

		if next.err != nil {
			panic("Error encountered. Call Has* funcs before calling this")
		} else {
			this.popBuff() // remove either empty or non-empty token

			if len(next.s) > 0 {
				// yeah! sure will return
				if v, e := converter(next.s); e == nil {
					return v
				} else {
					panic("Cannot convert to int: '" + next.s + "'")
				}
			}
		}
	}
	panic("should not reach here")
	return 0
}

// Reads a non-whitespace sequence of characters from the input.
// Leading whitespaces are skipped.
func (this *Scanner) Next() string {
	for {
		next := this.nextBuffedToken(nil).Value.(sd)

		if next.err != nil {
			panic("Error encountered. Call Has* funcs before calling this")
		} else {
			this.popBuff() // remove either empty or non-empty token

			if len(next.s) > 0 {
				return next.s
			}
		}
	}
	panic("should not reach here")
}

// Reads an int from the input. Call HasNextInt first to check whether an int can be read.
func (this *Scanner) NextInt() int {
	res := this.nextSth(func(s string) (v int64, e error) {
		v_, e := strconv.Atoi(s)
		v = int64(v_)
		return
	})
	return int(res)
}

// Reads an int64 from the input. Call HasNextInt64 first to check whether an int64 can be read.
func (this *Scanner) NextInt64() int64 {
	return this.nextSth(func(s string) (v int64, e error) {
		v, e = strconv.ParseInt(s, 10, 64)
		return
	})
}

// Reads a uint from the input. Call HasNextUint first to check whether an uint can be read.
func (this *Scanner) NextUint() uint {
	res := this.nextSth(func(s string) (v int64, e error) {
		v_, e := strconv.ParseUint(s, 10, 0)
		v = int64(v_)
		return
	})
	return uint(res)
}

// Reads a uint64 from the input. Call HasNextUint64 first to check whether an uint64 can be read.
func (this *Scanner) NextUint64() uint64 {
	res := this.nextSth(func(s string) (v int64, e error) {
		v_, e := strconv.ParseUint(s, 10, 64)
		v = int64(v_)
		return
	})
	return uint64(res)
}

// Reads the input stream from the current position until the line terminator is encountered, or
// the end of input stream is reached. The line terminator is not included.
// Call HasNextLine first to check whether a line can be read.
func (this *Scanner) NextLine() string {
	buf := bytes.NewBufferString("")

	for {
		next := this.nextBuffedToken(nil).Value.(sd)

		if next.err != nil {
			if buf.Len() == 0 {
				panic("Error encountered. Call Has* funcs before calling this")
			} else {
				// return last line
				return buf.String()
			}
		} else {
			this.popBuff()          // remove either empty or non-empty token
			buf.WriteString(next.s) // and put the string

			// is the delim a newline sign?
			if next.delim == '\n' {
				return buf.String()
			}

			// and the delim, too, because it was not a new line sign
			buf.WriteString(string(next.delim))
		}
	}

	panic("should not reach here")
	return ""
}

func (this *Scanner) hasNextSth(tester func(s string) bool) bool {
	after := (*list.Element)(nil)
	for {
		nextElement := this.nextBuffedToken(after)
		next := nextElement.Value.(sd)

		if next.err != nil {
			return false
		}

		if len(next.s) > 0 {
			// we have the data, check if it's an int/uint etc
			return tester(next.s)
		}

		// last was double-delimiter. so we go back to loop after skipping the first element.
		after = nextElement
	}
	panic("should not reach here")
}

// Checks if the input stream still has an int to be read using NextInt.
func (this *Scanner) HasNextInt() bool {
	return this.hasNextSth(func(s string) bool {
		_, e := strconv.Atoi(s)
		return e == nil
	})
}

// Checks if the input stream still has an int64 to be read using NextInt64.
func (this *Scanner) HasNextInt64() bool {
	return this.hasNextSth(func(s string) bool {
		_, e := strconv.ParseInt(s, 10, 64)
		return e == nil
	})
}

// Checks if the input stream still has a uint to be read using NextUint.
func (this *Scanner) HasNextUint() bool {
	return this.hasNextSth(func(s string) bool {
		_, e := strconv.ParseUint(s, 10, 0)
		return e == nil
	})
}

// Checks if the input stream still has a uint64 to be read using NextUint64.
func (this *Scanner) HasNextUint64() bool {
	return this.hasNextSth(func(s string) bool {
		_, e := strconv.ParseUint(s, 10, 64)
		return e == nil
	})
}

// Checks if the input stream still has a sequence of non-whitespace characters to be read using Next.
func (this *Scanner) HasNext() bool {
	return this.hasNextSth(func(string) bool { return true })
}

// Checks if the input stream still has a line to read using NextLine.
func (this *Scanner) HasNextLine() bool {
	// simple. Just check if next token is not EOF
	next := this.nextBuffedToken(nil).Value.(sd)
	return next.err == nil
}
