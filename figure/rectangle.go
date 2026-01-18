package figure

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) Figure {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r *Rectangle) Square() (float64, error) {
	return r.width * r.height, nil
}

func (r *Rectangle) Perimeter() (float64, error) {
	return 2 * (r.width + r.height), nil
}
