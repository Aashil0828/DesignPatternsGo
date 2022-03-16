package main

type HPAdapter struct {
	wireSize int
}

func (d *HPAdapter) getWireSize() int {
	return d.wireSize
}
