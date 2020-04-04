package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValue1(t *testing.T) {
  a:=100
  va,vp:=reflect.ValueOf(a),reflect.ValueOf(&a).Elem()
  fmt.Println(va.CanAddr(),va.CanSet())
  fmt.Println(vp.CanAddr(),vp.CanSet())
}

func TestValue2(t *testing.T) {
	type user struct {
		Name string
		Age int
	}

	u:=user{
		"q.yuhen",
		60,
	}

	v:=reflect.ValueOf(&u)

	if !v.CanInterface() {
		println("CanInterface:fail.")
		return
	}

	p,ok:=v.Interface().(*user)
	if !ok{
		println("Interface fail")
		return
	}

	p.Age++
	fmt.Printf("%+v\n",u)
}