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

	println(math.Pow(2, 10))
	println(math.Pow(2, 20))
	println(math.Pow(2, 30))

}

func TestConst3(t *testing.T) {
	const (
		a = iota
		b
		c, d = iota, iota //同一行值一样
		e    = iota
	)
	println(a, b, c, d, e)
}

func TestConst4(t *testing.T) {
	const (
		a = iota
		b
		c = 100
		d     //100,与上一行右值表达式相同
		e = iota
		f
	)
	println(a, b, c, d, e, f)
}
