package tasks

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x_ float64, y_ float64) *Point {
	return &Point{x: x_, y: y_}
}

func Distance(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func Run24() {
	A := NewPoint(3.0, 0)
	B := NewPoint(0, 4.0)
	fmt.Println("\tDistance =", Distance(A, B))
}
