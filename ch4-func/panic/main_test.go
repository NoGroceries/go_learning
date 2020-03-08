package main

import "testing"

func test(x,y int) {
	z:=0

	func() {
		defer func() {
			if recover()!=nil {
				z=0
			}
		}()
		z=x/y
	}()

	println("x/y=",z)
}

func TestT(t *testing.T) {
	test(5,0)
}
