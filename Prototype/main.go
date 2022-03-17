package main

func main() {
	l := Laptop{name: "HP", size: 10}
	b := Bag{name: "American Tourister", size: 10}
	lclone := l.clone()
	bclone := b.clone()
	lclone.print()
	bclone.print()

}
