package main

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

type Coordinates struct {
	x int
	y int
}

type Circle struct {
	Coordinates
	radius int
}

func NewRectangle(height int, width int) *Rectangle {
	return &Rectangle{height, width}
}

func (r Rectangle) Surface() int {
	return r.Height * r.Width
}

func main() {

	a := Rectangle{Height: 7}
	fmt.Println(a)

	r := NewRectangle(2, 3)
	fmt.Println(r)
	fmt.Println(&r)
	// *r-valor, &r-memoria
	fmt.Println(*r)

	c := Circle{Coordinates{1, 2}, 3}
	fmt.Printf("%+v\n", c)
	fmt.Println(c)
	fmt.Printf("%+v\n", c.Coordinates)
	fmt.Println(c.Coordinates)
	fmt.Println(c.x, c.y)

	rec := Rectangle{2, 3}
	fmt.Printf("rectangle %v has surface %d\n", rec, rec.Surface())

}
