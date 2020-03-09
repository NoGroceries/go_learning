package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	s := "茂宏"

	//byte 方式
	for i := 0; i < len(s); i++ {
		println(i,": ",s[i])

	}

	//rune方式
	for i,v:=range s {
		println(i,":",string(v))
	}

	//以下两种都可以得到str的字符串长度
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(s))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(s)))
}

//测试string拼接性能
func test1(num int)string {
	var s string
	for i:=0;i<num;i++ {
		s+="a"
	}
	return s
}

func test2(num int)string {
	s:=make([]string,num)
	for i:=0;i<num;i++ {
		s[i]="a"
	}
	return strings.Join(s,"")
}

func test3(num int) string{
	var b bytes.Buffer
	b.Grow(num)
	for i:=0;i<num;i++ {
		b.WriteString("a")
	}
	return b.String()
}

func BenchmarkTest(b *testing.B) {
	for i:=0;i<b.N;i++ {
		test1(10000)
	}
}

func BenchmarkTest2(b *testing.B) {
	for i:=0;i<b.N;i++ {
		test2(10000)
	}
}

func BenchmarkTest3(b *testing.B) {
	for i:=0;i<b.N;i++ {
		test3(10000)
	}
}
