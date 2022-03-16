package main

type DellMouse struct {
	mouseSize int
}

func (d *DellMouse) getMouseSize() int {
	return d.mouseSize
}
