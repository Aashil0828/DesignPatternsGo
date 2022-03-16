package main

type IndianHouseBuilder struct {
	wallType   string
	floorType  string
	windowType string
}

func (i *IndianHouseBuilder) setWallType() {
	i.wallType = "Cement & Bricks"
}
func (i *IndianHouseBuilder) setFloorType() {
	i.floorType = "Tiles"
}
func (i *IndianHouseBuilder) setWindowType() {
	i.windowType = "Wooden"
}
func (i *IndianHouseBuilder) getResult() *House {
	return &House{
		wallType:   i.wallType,
		floorType:  i.floorType,
		windowType: i.windowType,
	}
}
