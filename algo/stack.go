package main

import (
	"errors"
	"fmt"
)

type StackT struct {
	data []string
	size uint
	top  uint
}

func NewStack(max uint) *StackT {
	s := make([]string, max)
	return &StackT{data: s, size: max, top: 0}
}

func (s *StackT) Push(pushed string) error {
	n := s.top
	if n >= s.size-1 {
		return errors.New("Stack overflow")
	}
	s.top++
	s.data[n] = pushed
	return nil
}

func (s *StackT) Pop() (string, error) {
	n := s.top
	if n == 0 {
		return string(""), fmt.Errorf("Stack underflow")
	}
	top := s.data[n-1]
	s.top--
	return top, nil
}

func (s *StackT) Print() {
	n := s.top
	fmt.Println("Cap:", s.size, "Size:", n)
	var i uint
	for i = 0; i < n; i++ {
		fmt.Printf("\t%d:\t%s\n", i, s.data[i])
	}
}

func main() {
	stack := NewStack(10)
	stack.Print()
	stack.Push("boo")
	stack.Print()
	popped, _ := stack.Pop()
	fmt.Printf("Stack top is %s\n", popped)
	stack.Print()
	stack.Push("moo")
	stack.Push("zoo")
	stack.Print()
	popped2, _ := stack.Pop()
	fmt.Printf("Stack top is %s\n", popped2)
	stack.Print()
}
