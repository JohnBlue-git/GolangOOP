package services

import (
	"oop-example/models"
	"oop-example/utils"
	"fmt"
)

// ShapeCalculator provides calculation services for shapes
type ShapeCalculator struct {
	shapes []models.Shape
}

// NewShapeCalculator creates a new ShapeCalculator instance
func NewShapeCalculator() *ShapeCalculator {
	return &ShapeCalculator{
		shapes: make([]models.Shape, 0),
	}
}

// AddShape adds a shape to the calculator
func (sc *ShapeCalculator) AddShape(shape models.Shape) {
	sc.shapes = append(sc.shapes, shape)
}

// TotalArea calculates the total area of all shapes
func (sc *ShapeCalculator) TotalArea() float64 {
	areas := make([]float64, len(sc.shapes))
	for i, shape := range sc.shapes {
		areas[i] = shape.Area()
	}
	return utils.Sum(areas)
}

// TotalPerimeter calculates the total perimeter of all shapes
func (sc *ShapeCalculator) TotalPerimeter() float64 {
	perimeters := make([]float64, len(sc.shapes))
	for i, shape := range sc.shapes {
		perimeters[i] = shape.Perimeter()
	}
	return utils.Sum(perimeters)
}

// AverageArea calculates the average area of all shapes
func (sc *ShapeCalculator) AverageArea() float64 {
	areas := make([]float64, len(sc.shapes))
	for i, shape := range sc.shapes {
		areas[i] = shape.Area()
	}
	return utils.Average(areas)
}

// PrintShapeDetails prints detailed information about a shape
func PrintShapeDetails(shape models.Shape) {
	fmt.Printf("\n%s Details:\n", shape.Name())
	fmt.Printf("  Area:      %.2f\n", utils.Round(shape.Area(), 2))
	fmt.Printf("  Perimeter: %.2f\n", utils.Round(shape.Perimeter(), 2))
}

// PrintAllShapes prints information about all shapes in the calculator
func (sc *ShapeCalculator) PrintAllShapes() {
	fmt.Println("\n=== All Shapes ===")
	for i, shape := range sc.shapes {
		fmt.Printf("\nShape %d: %s\n", i+1, shape.Name())
		fmt.Printf("  Area:      %.2f\n", utils.Round(shape.Area(), 2))
		fmt.Printf("  Perimeter: %.2f\n", utils.Round(shape.Perimeter(), 2))
	}
}

// PrintSummary prints a summary of all calculations
func (sc *ShapeCalculator) PrintSummary() {
	fmt.Println("\n=== Summary ===")
	fmt.Printf("Total Shapes:      %d\n", len(sc.shapes))
	fmt.Printf("Total Area:        %.2f\n", utils.Round(sc.TotalArea(), 2))
	fmt.Printf("Total Perimeter:   %.2f\n", utils.Round(sc.TotalPerimeter(), 2))
	fmt.Printf("Average Area:      %.2f\n", utils.Round(sc.AverageArea(), 2))
}

// compareAreas is an unexported (private) helper function
func compareAreas(s1, s2 models.Shape) models.Shape {
	if s1.Area() > s2.Area() {
		return s1
	}
	return s2
}

// GetLargestShape returns the shape with the largest area
func (sc *ShapeCalculator) GetLargestShape() models.Shape {
	if len(sc.shapes) == 0 {
		return nil
	}
	
	largest := sc.shapes[0]
	for _, shape := range sc.shapes[1:] {
		largest = compareAreas(largest, shape)
	}
	return largest
}
