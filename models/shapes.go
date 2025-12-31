package models

import "math"

// Shape is an interface that defines behavior for geometric shapes
// Exported (public) because it starts with uppercase
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Rectangle represents a rectangular shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Name returns the shape name
func (r Rectangle) Name() string {
	return "Rectangle"
}

// Circle represents a circular shape
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the circumference of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Name returns the shape name
func (c Circle) Name() string {
	return "Circle"
}

// Triangle represents a triangular shape
type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
	SideC  float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of a triangle
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// Name returns the shape name
func (t Triangle) Name() string {
	return "Triangle"
}

// NewRectangle creates a new rectangle with validation
// This is a factory function (exported)
func NewRectangle(width, height float64) *Rectangle {
	if width <= 0 || height <= 0 {
		return nil
	}
	return &Rectangle{Width: width, Height: height}
}

// NewCircle creates a new circle with validation
func NewCircle(radius float64) *Circle {
	if radius <= 0 {
		return nil
	}
	return &Circle{Radius: radius}
}

// NewTriangle creates a new triangle with validation
func NewTriangle(base, height, sideA, sideB, sideC float64) *Triangle {
	if base <= 0 || height <= 0 || sideA <= 0 || sideB <= 0 || sideC <= 0 {
		return nil
	}
	return &Triangle{
		Base:   base,
		Height: height,
		SideA:  sideA,
		SideB:  sideB,
		SideC:  sideC,
	}
}
