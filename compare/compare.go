// Copyright 2011 Graham Woodward (grahamwoodward00@gmail.com).
//  All rights reserved.Use of this source code is governed by 
// the New BSD License that can be found in the LICENSING file.

//This package makes a user-defined type comparable provided that
//  such a type has a method with the signature
//  func ( st *sometype) Compare( *sametype interface{}) int
//  which returns 0 if the two items being compared are equal,
//  1 if the receiver is greater,and -1 if less than
  
package compare

type comparer interface {
		Compare( object interface{}) int
}
func IsEqual(first, second interface{}) bool {
	f,_ := first.(comparer)
	s,_ := second.(comparer)
	if f.Compare(s) == 0 {return true}
	return false
}
func IsNotEqual(first, second interface{}) bool {
	return ! IsEqual(first, second)
}

func IsGreaterThan(first,second interface{}) bool {
	f,_ := first.(comparer)
	s,_ := second.(comparer)
	if f.Compare(s) == 1 {return true}
	return false
}
func IsLessThan(first, second interface{}) bool {
	f,_ := first.(comparer)
	s,_ := second.(comparer)
	if f.Compare(s) == -1 {return true}
	return false
}
func IsGreaterThanOrEqual(first, second interface{}) bool {
	f,_ := first.(comparer)
	s,_ := second.(comparer)
	if (f.Compare(s) == 0) || (f.Compare(s) == 1) {return true}
	return false
}
func IsLessThanOrEqual(first, second interface{}) bool {
	f,_ := first.(comparer)
	s,_ := second.(comparer)
	if (f.Compare(s) == 0) || (f.Compare(s) == -1) {return true}
	return false
}	
