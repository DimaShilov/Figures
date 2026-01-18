package figure

import "math"

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Figure {
	return &Circle{radius}
}

func (c *Circle) Square() (float64, error) {
	return math.Pow(c.radius, 2) * math.Pi, nil
}
func (c *Circle) Perimeter() (float64, error) {
	return 2 * c.radius * math.Pi, nil
}
