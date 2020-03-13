package ch6_method

import (
	"fmt"
	"testing"
)

type N int

func (n N) test() {
	fmt.Printf("test.n: %p,%d\n",&n ,n)
}

func Test1(t *testing.T) {
	var n N = 25
	fmt.Printf("n:%p,%d\n",&n,n)

	f1:=N.test
	f1(n)

	f2:=(*N).test
	f2(&n)
}
