package main

import (
	"oop-example/models"
	"oop-example/services"
	"oop-example/utils"
	"fmt"
)

// Types for basic OOP example
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	fmt.Println("=== OOP and Package Design in Go ===")

	// Part 1: Basic OOP Example (original)
	fmt.Println("\n--- Part 1: Basic OOP ---")
	basicOOPExample()

	// Part 2: Package Design Example
	fmt.Println("\n--- Part 2: Package Design ---")
	packageDesignExample()
}

// basicOOPExample demonstrates basic OOP concepts with inline types
func basicOOPExample() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	printArea(r)
	printArea(c)
}

// packageDesignExample demonstrates how to use multiple packages
func packageDesignExample() {
	// Create shapes using factory functions from models package
	rect := models.NewRectangle(5, 10)
	circle := models.NewCircle(7)
	triangle := models.NewTriangle(6, 8, 6, 8, 10)

	// Demonstrate individual shape operations
	fmt.Println("\n--- Individual Shapes ---")
	if rect != nil {
		services.PrintShapeDetails(rect)
	}
	if circle != nil {
		services.PrintShapeDetails(circle)
	}
	if triangle != nil {
		services.PrintShapeDetails(triangle)
	}

	// Use the calculator service to manage multiple shapes
	fmt.Println("\n--- Using ShapeCalculator Service ---")
	calculator := services.NewShapeCalculator()

	// Add shapes to calculator
	if rect != nil {
		calculator.AddShape(rect)
	}
	if circle != nil {
		calculator.AddShape(circle)
	}
	if triangle != nil {
		calculator.AddShape(triangle)
	}

	// Print all shapes
	calculator.PrintAllShapes()

	// Print summary statistics
	calculator.PrintSummary()

	// Find and display largest shape
	largest := calculator.GetLargestShape()
	if largest != nil {
		fmt.Println("\n--- Largest Shape ---")
		fmt.Printf("The largest shape is: %s with area %.2f\n", 
			largest.Name(), utils.Round(largest.Area(), 2))
	}

	// Demonstrate utility functions
	fmt.Println("\n--- Utility Functions ---")
	values := []float64{10.5, 20.7, 15.3, 8.9}
	fmt.Printf("Values: %v\n", values)
	fmt.Printf("Sum: %.2f\n", utils.Sum(values))
	fmt.Printf("Average: %.2f\n", utils.Average(values))
	fmt.Printf("Max: %.2f\n", utils.Max(values[0], values[1]))
	fmt.Printf("Min: %.2f\n", utils.Min(values[0], values[1]))
	fmt.Printf("Rounded Pi: %.2f\n", utils.Round(3.14159265359, 2))

	// Demonstrate validation
	fmt.Println("\n--- Validation Example ---")
	testValue := 42.5
	fmt.Printf("Is %.2f positive? %v\n", testValue, utils.ValidatePositive(testValue))
	testValue = -10.0
	fmt.Printf("Is %.2f positive? %v\n", testValue, utils.ValidatePositive(testValue))

	fmt.Println("\n--- Package Design Benefits ---")
	fmt.Println("✓ models package: Encapsulates shape data structures and interfaces")
	fmt.Println("✓ utils package: Provides reusable utility functions")
	fmt.Println("✓ services package: Implements business logic and orchestration")
	fmt.Println("✓ main package: Demonstrates how packages work together")
	fmt.Println("✓ Clear separation of concerns and code organization")
}
