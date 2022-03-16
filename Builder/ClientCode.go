package main

import "fmt"

func main() {
	indianbuilder := getBuilder("Indian")
	director := Director{}
	director.setBuilder(indianbuilder)
	indianhouse := director.makeHouse()
	foreignbuilder := getBuilder("Foreign")
	director.setBuilder(foreignbuilder)
	foreignhouse := director.makeHouse()
	fmt.Println("Indian House")
	fmt.Printf("Wall Type: %v\n", indianhouse.wallType)
	fmt.Printf("Floor Type: %v\n", indianhouse.floorType)
	fmt.Printf("Window Type: %v\n", indianhouse.windowType)
	fmt.Println("Foreign House")
	fmt.Printf("Wall Type: %v\n", foreignhouse.wallType)
	fmt.Printf("Floor Type: %v\n", foreignhouse.floorType)
	fmt.Printf("Window Type: %v\n", foreignhouse.windowType)
}
