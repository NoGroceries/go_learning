package main

func main(){
	x := 100
	println(&x,x)
	{
		x,y:=200,300
		println(&x,x,y) //不同作用域，全部是新变量定义
	}
}
