package main

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {

	data :=[3]string{"a","b","c"}

	for i := range data {  //只返回1st value
		println(i)
	}

	for _,v:=range data { //忽略 1st value
		println(v)
	}

	for range data {

	}
}

func TestFor2(t *testing.T) {
	data := [3]int{10,20,30}

	for i,v:=range data {
		if i==0 {
			data[0] +=100
			data[1] +=200
			data[2] +=300
		}

		fmt.Printf("%d %d",v,data[i])
	}
	fmt.Printf("%v",data)
}

func TestFor3(t *testing.T) {
	data := [3]int{10,20,30}

	for i,v:=range data[:] {
		if i==0 {  //
			data[0] +=100
			data[1] +=200
			data[2] +=300
		}//第一轮执行后，变成了 d[0] 110  d[1] 220 d[2] 330

		fmt.Printf("%d: %d\n",v,data[i])
	}
	//fmt.Printf("%v",data)
}
