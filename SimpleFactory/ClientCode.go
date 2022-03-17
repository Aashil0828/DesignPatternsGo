package main

import "fmt"

func main() {
	americanbag := getBag("AT")
	skybag := getBag("SB")
	printdetails(americanbag)
	printdetails(skybag)
}

func printdetails(b iBag) {
	fmt.Println(b.getName())
	fmt.Printf("Size: %v\n", b.getSize())
	fmt.Printf("Chains: %v\n", b.getChains())
}
