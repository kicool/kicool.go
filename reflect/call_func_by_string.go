package main

import (
	"reflect"
	"fmt"
)

type MindMap struct {
	Title string
	Author string
}

func (mm MindMap) string() string {
	return fmt.Sprint("Title: ", mm.Title, "Author: ", mm.Author)
}

func main() {
	m := MindMap{"MindMap", "kicool"}
	fmt.Printf("%v %s", m, reflect.ValueOf(m))
}
