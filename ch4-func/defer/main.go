package main

func main() {
	//for i:=0;i<10;i++ {
	//	println(i)
	//	defer func() {
	//		println("defer")
	//	}()
	//}
	//println("main") //先执行main 后执行defer
	do:=func(n int) {
		println(n)
		defer func() {
			println("defer")
		}()
	}

	for i:=0;i<10 ;i++  {
		do(i)
	}
	println("main")
}
