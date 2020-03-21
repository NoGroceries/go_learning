package main

import (
	"reflect"
)

type X int

func main()  {
	var x X =10
	t:=reflect.TypeOf(x)
	println(t.Name(),t.Kind())
}

