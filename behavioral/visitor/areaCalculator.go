package main

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(square *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(circle *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(rectangle *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
