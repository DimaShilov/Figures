package triangle

import (
	"math"
)

func HalfTrianglePerimeter(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return p
}

func GeronTriangleSquare(a, b, c float64) float64 {
	p := HalfTrianglePerimeter(a, b, c)
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return s
}

func RecTriangleSquare(a, b float64) float64 {
	s := (a * b) / 2
	return s
}

func SinTriangleSquare(a, b, degree float64) float64 {
	s := (a * b * math.Sin(degree)) / 2
	return s
}
