package _map

import (
	"fmt"
	"testing"
)

//map 基本操作
func TestMap(t *testing.T) {
	m:=map[string]int {
		"a":1,
		"b":2,
	}
	m["b"] = 3
	m["c"] = 4

	if v,ok:=m["d"];ok{
		println(v)
	}

	delete(m,"d") //不存在，不会报错
	delete(m,"b")
	fmt.Printf("%v",m)
}

func TestMap2(t *testing.T) {
	var m1 map[string]int
	m2:=map[string]int{} //内容为空的字典与nil不同
	println(m1==nil,m2==nil)
}
