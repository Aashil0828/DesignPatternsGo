package main

type ForeignHouseBuilder struct {
	wallType   string
	floorType  string
	windowType string
}

func (i *ForeignHouseBuilder) setWallType() {
	i.wallType = "Wooden"
}
func (i *ForeignHouseBuilder) setFloorType() {
	i.floorType = "Carpet"
}
func (i *ForeignHouseBuilder) setWindowType() {
	i.windowType = "Glass"
}
func (i *ForeignHouseBuilder) getResult() *House {
	return &House{
		wallType:   i.wallType,
		floorType:  i.floorType,
		windowType: i.windowType,
	}
}
