package main

type DellAdapter struct {
	wireSize int
}

func (d *DellAdapter) getWireSize() int {
	return d.wireSize
}
