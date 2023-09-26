package main

import (
	"fmt"
	"math"
)

// Структура Point для представления точек в 2D-пространстве

type Point struct {
	x, y float64
}

// Конструктор для создания новой точки

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между двумя точками

func (p Point) DistanceTo(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки с использованием конструктора

	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками с использованием метода DistanceTo

	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
