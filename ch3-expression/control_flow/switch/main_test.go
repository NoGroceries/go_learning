package mian

import "testing"

func TestSwitch(t *testing.T) {
	a, b, c, d, x := 1, 2, 3, 4, 1

	switch x {
	case a:   //单条件，内容为空。隐式"case a:break;"
	case b:
		println("b")
	case c:
		println("c")
	case d:
		println("d")
	}
}

//fallthrough 贯通后续case
func TestSwitch2(t *testing.T) {
	x:=5
	switch x {
	case 5:  //case中可以写多条语句
		x +=10
		println(x) //15
		fallthrough  //继续执行下一个case，但不再匹配条件表达式
	case 6:
		x +=10
		println(x)//25
	}
}
