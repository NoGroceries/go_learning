package argument

import (
	"fmt"
	"testing"
)

func f(x,y int,s string,_ bool) *int{
	return nil
}

func TestF(t *testing.T) {
	f(1,2,"x",true)
}

func f2(x *int) {
	fmt.Printf("pointer:%p, target:%v\n",&x,x)
}

func TestF2(t *testing.T) {
	a:=0x100
	p:=&a
	fmt.Printf("pointer:%p, target:%v\n", &p,p)
	f2(p)
}

func TestF3(t *testing.T) {
	var strss= []string{
		"qwr",
		"234",
		"yui",

	}
	var strss2= []string{
		"qqq",
		"aaa",
		"zzz",
		"zzz",
	}
	strss=append(strss,strss2...) //strss2的元素被打散一个个append进strss
	fmt.Println(strss)

}
