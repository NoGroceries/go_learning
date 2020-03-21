package main

import (
	"reflect"
	"testing"
)

type X int

func TestExe1(t *testing.T) {
	var x X = 10
	d:=reflect.TypeOf(x)
	println(d.Kind())
}
