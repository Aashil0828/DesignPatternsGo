package main

type AmericanTourister struct {
	Bag
}

func newAmericanTourister() iBag {
	return &AmericanTourister{Bag{name: "American Tourister", size: 10, chains: 3}}
}
