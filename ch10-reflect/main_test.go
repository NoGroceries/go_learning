package main

import (
	"fmt"
	"reflect"
	"testing"
)

type X int

func TestExe1(t *testing.T) {
	var x X = 10
	a := reflect.TypeOf(x)
	println(a.Kind())
	fmt.Println(a.Name(), a.Kind())

	b := reflect.TypeOf(&x)
	fmt.Println(b.Name(), b.Kind())
	fmt.Println(b.Elem())
}

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func TestExe2(t *testing.T) {
	var m manager
	a := reflect.TypeOf(&m)

	if a.Kind() == reflect.Ptr {
		a = a.Elem()
	}

	for i := 0; i < a.NumField(); i++ {
		f := a.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}

}

//对于匿名字段，可用多级索引直接访问
func TestExe3(t *testing.T) {
	var m manager
	a := reflect.TypeOf(m)

	name, _ := a.FieldByName("name")
	fmt.Println(name.Name, name.Type)

	age := a.FieldByIndex([]int{0, 1})
	fmt.Println(age.Name, age.Type)

}

type A int
type B struct {
	A
}

func (A) av() {

}
func (*A) ap() {

}

func (B) bv() {

}
func (*B) bp() {

}
func TestExe4(x *testing.T) {
	var b B
	t := reflect.TypeOf(&b)

	s := []reflect.Type{t, t.Elem()} //slice

	for _, t := range s {
		fmt.Println(t, ":")

		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(" ", t.Method(i))
		}
	}
}
