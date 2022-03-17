package main

type iBag interface {
	setSize(size int)
	setChains(chains int)
	setName(name string)
	getSize() int
	getChains() int
	getName() string
}
