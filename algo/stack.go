package main

import (
	"errors"
	"fmt"
)

const max = 100

type StackT struct {
	stack [max]string
	size  int
}

func (s *StackT) push(pushed string) error {
	n := s.size
	if n >= max-1 {
		return errors.New("Stack overflow")
	}
	s.size++
	s.stack[n] = pushed
	return nil
}

func (s *StackT) pop() (string, error) {
	n := s.size
	if n == 0 {
		return string(""), errors.New("Stack underflow")
	}
	top := s.stack[n-1]
	s.size--
	return top, nil
}

func (s *StackT) print_all() {
	n := s.size
	fmt.Println("Size:", n)
	for i := 0; i < n; i++ {
		fmt.Printf("\t%d:\t%s\n", i, s.stack[i])
	}
}

func main() {
	stack := new(StackT)
	stack.print_all()
	stack.push("boo")
	stack.print_all()
	popped, _ := stack.pop()
	fmt.Printf("Stack top is %s\n", popped)
	stack.print_all()
	stack.push("moo")
	stack.push("zoo")
	stack.print_all()
	popped2, _ := stack.pop()
	fmt.Printf("Stack top is %s\n", popped2)
	stack.print_all()
}
