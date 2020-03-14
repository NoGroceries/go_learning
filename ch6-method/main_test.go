package ch6_method

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type N int

func (n N) test() {
	fmt.Printf("test.n: %p,%d\n", &n, n)
}

func Test1(t *testing.T) {
	var n N = 25
	fmt.Printf("n:%p,%d\n", &n, n)

	f1 := N.test
	f1(n)

	f2 := (*N).test
	f2(&n)
}

func TestFile(t *testing.T) {

	file, err := os.Open("/null")
	defer file.Close()
	if err != nil {
		fmt.Println("close error: ", err)

	} else {
		fmt.Println("close no error")
	}

}


func lastLen(s string)int {
	if !strings.Contains(s," ") {
		return len(s)
	}
	ss:=strings.Split(s," ")
	return len(ss[len(ss)-1])

}

