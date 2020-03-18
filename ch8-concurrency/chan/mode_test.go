package _chan

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			println("recv: ", x)
		}
	}()
	return r
}

//使用工厂方法将goroutine与chan绑定
func Test1(t *testing.T) {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	close(r.data)
	r.Wait()
}

//chan 用在pool
type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p:
	default:
		v = make([]byte, 10)
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p<-b:
	default:

	}
}

func Test2(t *testing.T) {
	p :=newPool(10)
	var b =[]byte{1,3}
	p.put(b)
	a:=p.get()
	println(len(a)," ",cap(a))
	fmt.Printf("%T",a)
}

//timeout和tick chan实现
func Test3(t *testing.T) {
	go func() {
		for{
			select {
			case <-time.After(time.Second*5):
				fmt.Println("timeout ...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick:=time.Tick(time.Second)
		for{
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct {})(nil)//堵塞进程
}
