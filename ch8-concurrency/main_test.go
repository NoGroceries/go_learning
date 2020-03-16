package main

import (
	"fmt"
	"strings"
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


func compressString(S string) string {
	var newStr strings.Builder
	length :=len(S)
	for i:=0;i<length; {
		count:=1
		if length == 1 {
			return S
		}
		if i==length-1 &&S[length-1] != S[length-2] {
			newStr.WriteString(string(S[i]) + fmt.Sprintf("%d",count))
			break
		}
		v :=S[i]
		j :=i+1
		for ;j<length;j++ {
			if S[j]==v {
				count++
				if j == length-1 {
					newStr.WriteString(string(v) + fmt.Sprintf("%d",count))
					break;
				}
				continue
			}else {

				newStr.WriteString(string(v) + fmt.Sprintf("%d",count))
				break;
			}

		}
		i=j
	}



	s2:=newStr.String()

	if len(s2)>=length {
		return S
	}else {
		return s2
	}

}

func Test11(t *testing.T) {
	println(compressString("a"))
}
