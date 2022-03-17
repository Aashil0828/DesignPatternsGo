package main

import "fmt"

func getBag(bagname string) iBag {
	if bagname == "AT" {
		return newAmericanTourister()
	} else if bagname == "SB" {
		return newSkyBag()
	} else {
		fmt.Println("Enter valid Bag name")
		return nil
	}
}
