package shapes

type Square struct {
	Side float64
}

func NewSquare(side float64) Square {
	return Square{
		Side: side,
	}
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (s Square) Name() string {
	return "Square"
}

func (s Square) GetSide() float64 {
	return s.Side
}

func (s Square) GetDiagonal() float64 {
	return s.Side * 1.414213562373095
}
