package main

import ()

type T struct {
	a int
	b int32
	c int64
}

/*
 * new(T) returns type *T while make(T) returns type T.
 */

func main() {
	var t *T = new(T)
	t.a = 10
	println(t)
	println(t.a)
	println(t.b)
	println(t.c)

	t1 := new(T)
	println(t1)

	m := make(map[string]int)
	println(m)
}
