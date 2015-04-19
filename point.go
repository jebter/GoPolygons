package GoPolygons

import (
	"math"
)

//Two-dimensional point
type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{X: x, Y: y}
}

//Comparison is the same point
func (pt Point) IsSame(p Point) bool {
	return pt.X == p.X && pt.Y == p.Y
}

//Get the distance between two points
func (pt Point) GetDistance(p Point) float64 {
	x := p.X - pt.X
	y := p.Y - pt.Y
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

//Get vector angle from pt to p, unit: degree
func (pt Point) GetBearing(p Point) float64 {
	x := pt.X - p.X
	y := pt.Y - p.Y
	return math.Atan2(y, x) * 180.0 / math.Pi
}
