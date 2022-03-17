package main

type Bag struct {
	size   int
	chains int
	name   string
}

func (b *Bag) setSize(size int) {
	b.size = size
}

func (b *Bag) getSize() int {
	return b.size
}
func (b *Bag) setChains(chains int) {
	b.chains = chains
}

func (b *Bag) getChains() int {
	return b.chains
}
func (b *Bag) setName(name string) {
	b.name = name
}

func (b *Bag) getName() string {
	return b.name
}
