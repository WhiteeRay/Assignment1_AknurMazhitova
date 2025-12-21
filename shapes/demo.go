package shapes

func DemonstrateShapes() {
	manager := NewShapeManager()

	rect := NewRectangle(5.0, 3.0)
	manager.AddShape(rect)

	circle := NewCircle(4.0)
	manager.AddShape(circle)

	square := NewSquare(6.0)
	manager.AddShape(square)

	triangle := NewTriangle(3.0, 4.0, 5.0)
	manager.AddShape(triangle)

	manager.AddShape(NewRectangle(8.0, 2.5))
	manager.AddShape(NewCircle(2.5))
	manager.AddShape(NewSquare(4.5))

	manager.PrintAllShapes()
}
