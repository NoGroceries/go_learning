package main

//消费者
func consumer(data chan int, done chan bool) {
	for x:=range data {
		println("recv",":",x)
	}
	done <-true
}

//生产者
func producer(data chan int) {
	for i:=0;i<4;i++ {
		data <-i
	}
	close(data)
}


func main() {
	done := make(chan bool)
	data:=make(chan int)

	go consumer(data,done)
	go producer(data)
	<-done //阻塞，直到等到消费者发回的结束信息，如果将消费者中done <-true注释掉
	       //将会报错 fatal error: all goroutines are asleep - deadlock!
}
