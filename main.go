package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func (p Person) shout() {
	fmt.Printf("%v %v !!\n", p.FirstName, p.LastName)
}

func main() {
	p := Person{"Jeng", "Mahapak"}
	p.shout()
}