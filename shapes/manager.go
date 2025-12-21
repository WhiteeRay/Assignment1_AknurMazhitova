package shapes

import (
	"fmt"
	"strings"
)

type ShapeManager struct {
	shapes []Shape
}

func NewShapeManager() *ShapeManager {
	return &ShapeManager{
		shapes: make([]Shape, 0),
	}
}

func (sm *ShapeManager) AddShape(shape Shape) {
	sm.shapes = append(sm.shapes, shape)
}

func (sm *ShapeManager) GetShapes() []Shape {
	return sm.shapes
}

func (sm *ShapeManager) GetTotalArea() float64 {
	total := 0.0
	for _, shape := range sm.shapes {
		total += shape.Area()
	}
	return total
}

func (sm *ShapeManager) GetTotalPerimeter() float64 {
	total := 0.0
	for _, shape := range sm.shapes {
		total += shape.Perimeter()
	}
	return total
}

func (sm *ShapeManager) GetShapeCount() int {
	return len(sm.shapes)
}

func (sm *ShapeManager) GetAverageArea() float64 {
	if len(sm.shapes) == 0 {
		return 0
	}
	return sm.GetTotalArea() / float64(len(sm.shapes))
}

func (sm *ShapeManager) GetLargestShape() (Shape, float64) {
	if len(sm.shapes) == 0 {
		return nil, 0
	}

	largest := sm.shapes[0]
	largestArea := largest.Area()

	for _, shape := range sm.shapes {
		area := shape.Area()
		if area > largestArea {
			largest = shape
			largestArea = area
		}
	}

	return largest, largestArea
}

func (sm *ShapeManager) GetSmallestShape() (Shape, float64) {
	if len(sm.shapes) == 0 {
		return nil, 0
	}

	smallest := sm.shapes[0]
	smallestArea := smallest.Area()

	for _, shape := range sm.shapes {
		area := shape.Area()
		if area < smallestArea {
			smallest = shape
			smallestArea = area
		}
	}

	return smallest, smallestArea
}

func PrintShapeDetails(shape Shape) {
	fmt.Printf("%-12s | Area: %8.2f | Perimeter: %8.2f\n",
		shape.Name(), shape.Area(), shape.Perimeter())
}

func (sm *ShapeManager) PrintAllShapes() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("=== Shape Collection ===")
	fmt.Println(strings.Repeat("=", 60))

	if len(sm.shapes) == 0 {
		fmt.Println("No shapes in the collection.")
		return
	}

	for i, shape := range sm.shapes {
		fmt.Printf("%d. ", i+1)
		PrintShapeDetails(shape)
	}

	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Total Shapes: %d\n", sm.GetShapeCount())
	fmt.Printf("Total Area: %.2f\n", sm.GetTotalArea())
	fmt.Printf("Total Perimeter: %.2f\n", sm.GetTotalPerimeter())
	fmt.Printf("Average Area: %.2f\n", sm.GetAverageArea())

	if largest, area := sm.GetLargestShape(); largest != nil {
		fmt.Printf("Largest Shape: %s (Area: %.2f)\n", largest.Name(), area)
	}

	if smallest, area := sm.GetSmallestShape(); smallest != nil {
		fmt.Printf("Smallest Shape: %s (Area: %.2f)\n", smallest.Name(), area)
	}
	fmt.Println()
}
