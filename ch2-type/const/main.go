package main

import "fmt"

func main() {
   const x = 123
   fmt.Printf("%T %d",x,x)

   const y = 1.23 //未使用，不会引发编译错误
}
