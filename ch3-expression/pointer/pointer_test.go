package pointer

import "testing"

//指针不能做加减法运算和类型转换
func TestPointer(t *testing.T) {
	x:=10
	p:=&x

	//p++
	var p2 *int = p

	p2=&x
	println(p == p2)

}

func TestPointer2(t *testing.T) {
	a:= struct {
		x int
	}{}

	a.x = 100
	p:=&a
	p.x+=100
	println(p.x)
}
