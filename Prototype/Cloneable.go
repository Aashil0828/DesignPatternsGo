package main

type Cloneable interface {
	clone() Cloneable
	print()
}
