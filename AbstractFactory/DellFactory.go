package main

type DellFactory struct {
}

func (d *DellFactory) createAdapter() Adapter {
	return &DellAdapter{3}
}

func (d *DellFactory) createMouse() Mouse {
	return &DellMouse{5}
}

func (d *DellFactory) createDesktop() Desktop {
	return &DellDesktop{10}
}
