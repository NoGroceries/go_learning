package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//map 基本操作
func TestMap(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	m["b"] = 3
	m["c"] = 4

	if v, ok := m["d"]; ok {
		println(v)
	}

	delete(m, "d") //不存在，不会报错
	delete(m, "b")
	fmt.Printf("%v", m)
}

func TestMap2(t *testing.T) {
	var m1 map[string]int
	m2 := map[string]int{} //内容为空的字典与nil不同
	println(m1 == nil, m2 == nil)
}

func TestMap3(t *testing.T) {
	m := make(map[string]int)

	go func() {
		for {
			m["a"] = 1
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}
}

func TestMap4(t *testing.T) {
	m:=make(map[string]int)
	var lock sync.RWMutex

	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			lock.Unlock()
			time.Sleep(time.Microsecond)

		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["b"]
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()

	time.Sleep(time.Second)
}
