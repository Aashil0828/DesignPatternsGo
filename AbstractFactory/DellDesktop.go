package main

type DellDesktop struct {
	screenSize int
}

func (d *DellDesktop) getScreenSize() int {
	return d.screenSize
}
