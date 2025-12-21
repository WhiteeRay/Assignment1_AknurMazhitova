package shapes

import "math"

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func NewTriangle(sideA, sideB, sideC float64) Triangle {
	return Triangle{
		SideA: sideA,
		SideB: sideB,
		SideC: sideC,
	}
}

func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t Triangle) Name() string {
	return "Triangle"
}

func (t Triangle) GetSides() (float64, float64, float64) {
	return t.SideA, t.SideB, t.SideC
}

func (t Triangle) IsValid() bool {
	return (t.SideA+t.SideB > t.SideC) &&
		(t.SideA+t.SideC > t.SideB) &&
		(t.SideB+t.SideC > t.SideA)
}

func (t Triangle) IsEquilateral() bool {
	return t.SideA == t.SideB && t.SideB == t.SideC
}

func (t Triangle) IsIsosceles() bool {
	return t.SideA == t.SideB || t.SideB == t.SideC || t.SideA == t.SideC
}
