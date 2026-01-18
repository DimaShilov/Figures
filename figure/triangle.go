package figure

import (
	"Figures/service/triangle"
)

type Triangle struct {
	base      float64
	leftSide  float64
	rightSide float64
}

func NewTriangle(a, b, c float64) Figure {
	return &Triangle{
		base:      a,
		leftSide:  b,
		rightSide: c,
	}
}

func (t *Triangle) Square() (float64, error) {
	resSquare := triangle.GeronTriangleSquare(t.base, t.leftSide, t.rightSide)
	return resSquare, nil
}
func (t *Triangle) Perimeter() (float64, error) {
	resPerimeter := 2 * triangle.HalfTrianglePerimeter(t.base, t.leftSide, t.rightSide)
	return resPerimeter, nil
}
