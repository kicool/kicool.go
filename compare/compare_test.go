// Copyright 2011 Graham Woodward (grahamwoodward00@gmail.com).
//  All rights reserved.Use of this source code is governed by 
// the New BSD License that can be found in the LICENSING file.

package compare

import ("testing" ; "fmt")

type TestObject struct {
		number int
}
func (to *TestObject) Compare(object interface {}) int {
	param,_ := object.(*TestObject)
	if to.number > param.number { return 1}
	if to.number < param.number { return -1}	
	return 0
}
var a,b,c,d = &TestObject{5},&TestObject{5},&TestObject{6},&TestObject{4}

func TestIsEqual(t *testing.T) {
	fmt.Println("Testing IsEqual....")
	if IsEqual(a,b) != true {t.Fail()}
}
func TestIsNotEqual(t *testing.T) {
	fmt.Println("Testing IsNotEqual....")
	if IsNotEqual(b,c) != true {t.Fail()}
}
func TestIsGreaterThan(t *testing.T) {
	fmt.Println("Testing IsGreaterThan....")
	if IsGreaterThan(c,b) != true {t.Fail()}
}
func TestIsLessThan(t *testing.T) {
	fmt.Println("Testing IsLessThan....")
	if IsLessThan(b,c) != true {t.Fail()}
}
func TestIsGreaterThanOrEqual(t *testing.T) {
	fmt.Println("Testing IsGreaterThanOrEqual(Greater)....")
	if IsGreaterThanOrEqual(c,b) != true {t.Fail()}
	fmt.Println("Testing IsGreaterThanOrEqual(Equal)....")
	if IsGreaterThanOrEqual(a,b) != true {t.Fail()}
}
func TestIsLessThanOrEqual(t *testing.T) {
	fmt.Println("Testing IsLessThanOrEqual(Greater)....")
	if IsLessThanOrEqual(b,c) != true {t.Fail()}
	fmt.Println("Testing IsLessThanOrEqual(Equal)....")
	if IsLessThanOrEqual(a,b) != true {t.Fail()}
}
