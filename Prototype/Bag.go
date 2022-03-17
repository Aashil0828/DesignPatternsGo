package main

import "fmt"

type Bag struct {
	name string
	size int
}

func (b *Bag) clone() Cloneable {
	return &Bag{name: b.name + "_clone", size: b.size}
}
func (l *Bag) print() {
	fmt.Println(l.name)
}
