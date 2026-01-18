package figure

type Figure interface {
	Square() (float64, error)
	Perimeter() (float64, error)
}
