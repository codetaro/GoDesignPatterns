package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (m *MiddleCoordinates) visitForSquare(square *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (m *MiddleCoordinates) visitForCircle(circle *Circle) {
	//TODO implement me
	fmt.Println("Calculating middle point coordinates for circle")
}

func (m *MiddleCoordinates) visitForRectangle(rectangle *Rectangle) {
	//TODO implement me
	fmt.Println("Calculating middle point coordinates for rectangle")
}
