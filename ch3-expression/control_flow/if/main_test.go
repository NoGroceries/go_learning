package main

import "testing"

func TestIf(t *testing.T) {
	x := 3

	if (x > 5) {  //可省略括号
		println("a")
	} else if x < 5 && x > 0 {
		println("b")
	} else {
		println("c")
	}
}
