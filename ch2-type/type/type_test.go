package _type

import "testing"

func TestType(t *testing.T){

	//1、定义结构体
	type Person struct {
		name string
		age byte
	}

	//2、自定义类型
	type myStr string

	//3、 定义接口
	type Personer interface {
        Run()
	}

	//4、定义函数
	type add func(a,b int) int


}
