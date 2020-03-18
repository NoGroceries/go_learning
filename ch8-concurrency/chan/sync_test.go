package _chan

import (
	"sync"
	"testing"
	"time"
)

//将mutex作为匿名字段时，相关方法必须实现为pointer-receiver，
//否则会因复制导致锁失效
type data struct {
	sync.Mutex
}

func(d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i:=0;i<5;i++ {
		println(s,i)
		time.Sleep(time.Second)
	}
}
func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	//var d =&data{}
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}