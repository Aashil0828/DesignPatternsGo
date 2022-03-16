package main

type HPDesktop struct {
	screenSize int
}

func (d *HPDesktop) getScreenSize() int {
	return d.screenSize
}
