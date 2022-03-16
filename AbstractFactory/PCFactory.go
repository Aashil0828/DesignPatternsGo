package main

import "fmt"

type PCFactory interface {
	createDesktop() Desktop
	createMouse() Mouse
	createAdapter() Adapter
}

func getFactory(company string) PCFactory {
	if company == "Dell" {
		return &DellFactory{}
	} else if company == "HP" {
		return &HPFactory{}
	} else {
		fmt.Println("Enter a valid Company")
		return nil
	}
}
