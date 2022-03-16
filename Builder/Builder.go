package main

import "fmt"

type Builder interface {
	setWallType()
	setFloorType()
	setWindowType()
	getResult() *House
}

func getBuilder(housetype string) Builder {
	if housetype == "Indian" {
		return &IndianHouseBuilder{}
	} else if housetype == "Foreign" {
		return &ForeignHouseBuilder{}
	} else {
		fmt.Println("Enter a valid House type")
		return nil
	}
}
