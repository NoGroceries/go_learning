package _continue

import "testing"

func TestContinue(t *testing.T) {
	outer:
		for x:=0;x<5;x++{
			for y:=0;y<10;y++{
				if y>2{
					println()
					continue outer //新用法
				}

				if x>2{
					break outer  //新用法
				}

				println(x,":",y," ")
			}
		}
}
