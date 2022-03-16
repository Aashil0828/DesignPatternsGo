package main

type HPMouse struct {
	mouseSize int
}

func (d *HPMouse) getMouseSize() int {
	return d.mouseSize
}
