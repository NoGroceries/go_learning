//slice实现stack
package main

import (
	"errors"
	"fmt"
)


func main() {
	stack :=make([]int,0,5)
	push:=func(i int)error{
		n:=len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}else {
			stack =stack[:n+1]
			stack[n] = i
		}
		return nil
	}


	pop:=func()(int,error){
		n:=len(stack)
		if n == 0 {
			return 0,errors.New("stack is empty")
		}
		x:=stack[n-1]
		stack =stack[:n-1]
		return x,nil
	}

	for i:=0;i<6;i++ {
		fmt.Printf("%v,%v\n",push(i),stack)
	}

	for i:=0;i<6;i++ {
		i,err:=pop()
		fmt.Printf("%d,%v,%v\n",i,stack,err)
	}

}