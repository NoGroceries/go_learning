package _chan

import "testing"

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
	<-done  //阻塞，直到close(done)
}
