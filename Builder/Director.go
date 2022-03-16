package main

type Director struct {
	builder Builder
}

func (d *Director) makeHouse() *House {
	d.builder.setFloorType()
	d.builder.setWallType()
	d.builder.setWindowType()
	return d.builder.getResult()
}
func (d *Director) setBuilder(b Builder) {
	d.builder = b
}
