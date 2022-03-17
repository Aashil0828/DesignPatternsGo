package main

type SkyBag struct {
	Bag
}

func newSkyBag() iBag {
	return &SkyBag{Bag{name: "Sky Bag", size: 8, chains: 2}}
}
