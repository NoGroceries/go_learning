package main

import "testing"

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
