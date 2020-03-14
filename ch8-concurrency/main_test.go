package ch8_concurrency

import (
	"sync"
	"testing"
	"time"
)

func TestConc(t *testing.T) {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")
		close(exit)
	}()

	println("main...")
	<-exit
	println("main exit.")
}

func TestConc2(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			println("goroutine done.")
		}()
	}
	println("main..")
	wg.Wait()
	println("main exit.")
}
