package dump

import (
	//	. "dump"
	"fmt"
	"go/parser"
	"go/token"
	"testing"
)

var empty = ""

type S struct {
	A int
	B int
}

type T struct {
	S
	C int
}

type Circular struct {
	c *Circular
}

func TestSimple0(t *testing.T) {
	Dump(1)
}

func TestSimpleArray(t *testing.T) {
	Dump([...]int{0, 1, 2, -1})
	Dump([5]int{-1, 0, 1, 2, -1})
}

func TestSimpleSlice(t *testing.T) {
	s := []int{0, 1, 2, 7, 100}
	s1 := s[:3]
	Dump(s)
	Dump(s1)
}

func TestSimpleMap(t *testing.T) {
	Dump(map[string]int{"satu": 1, "dua": 2})
}

func TestSimplePtr(t *testing.T) {
	i := 0
	Dump(&i)
	//Dump(t)
	Dump(&empty)
	Dump(&[][]int{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}})
}
func TestSimpleString(t *testing.T) {
	Dump(token.STRING)
}

func TestSimple(t *testing.T) {

	Dump(T{S{1, 2}, 3})

	bulet := make([]Circular, 3)
	bulet[0].c = &bulet[1]
	bulet[1].c = &bulet[2]
	bulet[2].c = &bulet[0]

	Dump(struct{ a []Circular }{bulet})
}

func TestDumpAST(t *testing.T) {
	// empty

	// func ParseFile(filename string, src interface{}, scope *ast.Scope, mode uint) (*ast.File, os.Error)
	fset := token.NewFileSet()
	file, e := parser.ParseFile(fset, "dump_test.go", nil, parser.ParseComments)
	if e != nil {
		fmt.Println("error", e)
	} else {
		//fmt.Printf("%#v\n", file);
		Dump(file)
	}
}
