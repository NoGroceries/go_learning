package main

import (
	"math"
	"testing"
)

func TestConst1(t *testing.T) {
	const (
		x = iota
		y
		z
	)
	println(x, y, z)
}

func TestConst2(t *testing.T) {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
	)
	println(KB, MB, GB)

	println(math.Pow(2,10))
	println(math.Pow(2,20))
	println(math.Pow(2,30))

}
