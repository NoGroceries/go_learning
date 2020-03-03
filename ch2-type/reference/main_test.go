package reference

import (
	"fmt"
	"testing"
)

//引用类型有 slice、map、channel

func TestMakeSlice(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 100)
	fmt.Println(s)
}

func TestMap(t *testing.T)  {
	m := make(map[string]int)
	m["a"] = 1
	fmt.Print(m)
}


