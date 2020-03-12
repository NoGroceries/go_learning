package _struct

import (
	"fmt"
	"testing"
)

func TestStruct (t *testing.T) {
	type file struct {
		name string
		attr struct{
			owner int
			perm int
		}
	}

	f:=file{
		name:"test.dat",
	//	attr: {owner:1 , perm: 1},
	}
	f.attr.owner=1
	f.attr.perm=1
	fmt.Printf("%v\n",f)
}
