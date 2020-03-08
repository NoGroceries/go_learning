package main

import "log"

//固定写法
func main() {
	defer func() {
		if err:=recover();err!=nil {
			log.Fatalln(err)
		}
	}()

	panic("i am dead")
	println("i")
}
