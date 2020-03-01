package main

import "fmt"

type user struct {
	name string
	age byte
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name="tom"
	m.age = 30
	m.title= "CTO"
	println(m.ToString())
}
