package main

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	const v = 10 //无显示类型声明的常量
	fmt.Printf("%T\n", v)
	var a byte = 8
	b := a + v //v 自动转换为byte/uint类型
	fmt.Printf("%T,%v\n", b, b)
	//fmt.Sprintf() func Sprintf(format string, a ...interface{}) string
	//返回string

	var c float32 = 3.2
	d := c + v
	fmt.Printf("%T,%v\n", d, d)
}

func TestBitOperation(t *testing.T) {
	b :=23
	x:=1<<b //shift count type int, must be unsigned integer
	println(x)

}

