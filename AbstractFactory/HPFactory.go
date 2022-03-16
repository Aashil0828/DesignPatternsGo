package main

type HPFactory struct {
}

func (d *HPFactory) createAdapter() Adapter {
	return &HPAdapter{8}
}

func (d *HPFactory) createMouse() Mouse {
	return &HPMouse{6}
}

func (d *HPFactory) createDesktop() Desktop {
	return &HPDesktop{12}
}
