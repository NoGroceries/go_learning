package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var a,b [2]int
	println(a == b)

	c:=[2]int{1,2}
	d:=[2]int{0,1}
	e:=[...]int{0,1}
	println(c == d)
	println(d == e)

}

func test(x *[2]int) {
	fmt.Printf("x:%p ,%v\n",x ,*x)
	x[1] +=100
}

func TestArrayCopy(t *testing.T) {
	a:=[2]int{10,20}
	test(&a)
	fmt.Printf("x:%p,%v\n",&a,a)
}