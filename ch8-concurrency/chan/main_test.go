package _chan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-c
			if !ok {
				return
			}
			println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	//c <- 4  //panic
	fmt.Printf("after close chan %d\n",<-c) //从已关闭接收数据，返回已缓冲数据或零值
	<-done  //
}


func TestChan2(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		  for v:=range c { //chan 支持range
		  	println(v)
		  }
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	//c <- 4  //panic
	<-done  //
}

//广播通知
func TestChan3(t *testing.T) {
	//var wg sync.WaitGroup
	ready:=make(chan struct{})

	for i:=0;i<3;i++{
		//wg.Add(1)

		go func(id int) {
			//defer wg.Done()
			println(id,": ready")
			<-ready
			println(id,": running...")
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready? Go!")
	close(ready)
	//wg.Wait()
}

//单向通道
func TestChan4(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	c:=make(chan int)
	var send chan<- int = c
	var recv <-chan int =c

	go func() {
		defer wg.Done()
		//close(recv)  //close不能用于接收端
		for x:=range recv {
			println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(send)
		for i:=0;i<3;i++ {
			send<-i
		}
	}()

	wg.Wait()
}
