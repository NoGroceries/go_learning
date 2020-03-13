package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	x :=[...]int{0,1,2,3,4,5,6,7,8,9}
	a:=x[2:5:7] //x[low:high:max]
	fmt.Printf("%v\n",a)
	println(len(a))  //len=high-low
	println(cap(a))//cap=max-low
	fmt.Printf("%T\n", a)
}

func array() [1024]int {
	var x=[1024]int{}
	for i:=0;i<1024;i++ {
		x[i]=i
	}
	return x
}

func slice()[]int {
	var x=make([]int,1024)
	for i:=0;i<1024;i++ {
		x[i]=i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i:=0;i<b.N;i++ {
		array()
	}
}
func BenchmarkSlice(b *testing.B) {
	for i:=0;i<b.N;i++ {
		slice()
	}
}

func TestCopy(t *testing.T) {
	s:=[]int{0,1,2,3,4,5,6,7,8,9}

	s1:=s[5:8] //5,6,7
    n:=copy(s[4:],s1)	//s[4:] {4,5,6,7,8,9}   对应位置copy，s1的3个元素替换s[4:]前3个元素
    fmt.Printf("%d,%v\n",n,s)


}
