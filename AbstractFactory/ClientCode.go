package main

import "fmt"

func main() {
	factory := getFactory("Dell")
	dellmouse := factory.createMouse()
	delldesktop := factory.createDesktop()
	delladapter := factory.createAdapter()
	factory2 := getFactory("HP")
	hpmouse := factory2.createMouse()
	hpdesktop := factory2.createDesktop()
	hpadapter := factory2.createAdapter()
	fmt.Println("Dell")
	fmt.Printf("Desktop size: %v\n", delldesktop.getScreenSize())
	fmt.Printf("Mouse size: %v\n", dellmouse.getMouseSize())
	fmt.Printf("Adapter Wire size: %v\n", delladapter.getWireSize())
	fmt.Println("HP")
	fmt.Printf("Desktop size: %v\n", hpdesktop.getScreenSize())
	fmt.Printf("Mouse size: %v\n", hpmouse.getMouseSize())
	fmt.Printf("Adapter Wire size: %v\n", hpadapter.getWireSize())

}
