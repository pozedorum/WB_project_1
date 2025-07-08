package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Distance(other Point) float64 {
	xc := math.Abs(p.x - other.x)
	yc := math.Abs(p.y - other.y)
	return math.Hypot(xc, yc)
}

func main() {
	p1 := NewPoint(3, 9)
	p2 := NewPoint(4, 1)
	fmt.Printf("Расстояние от первой до второй точки: %.2f\n", p1.Distance(*p2))
	fmt.Printf("Расстояние от второй точки до первой: %.2f\n", p2.Distance(*p1))
}
