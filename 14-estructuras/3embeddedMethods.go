package main

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) Surface() int {
	return r.Height * r.Width
}

type Box struct {
	Rectangle
	depth int
}

func (b Box) Volume() int {
	// return b.Surface() * b.depth
	return b.Rectangle.Surface() * b.depth
}

func main() {
	fmt.Println("Example Embedded methods")

	b := Box{Rectangle{10, 5}, 10}
	fmt.Printf("%+v\n", b)
	fmt.Println("Volume", b.Volume())
}
