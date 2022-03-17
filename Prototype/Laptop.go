package main

import "fmt"

type Laptop struct {
	name string
	size int
}

func (l *Laptop) clone() Cloneable {
	return &Laptop{name: l.name + "_clone", size: l.size}
}

func (l *Laptop) print() {
	fmt.Println(l.name)
}
