package GoPolygons

import (
	"math"
)

type Line struct {
	PointA Point
	PointB Point
}

type Rect struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

func NewLineByPoint(ptA Point, ptB Point) Line {
	return Line{PointA: ptA, PointB: ptB}
}

func NewLine(x1 float64, y1 float64, x2 float64, y2 float64) Line {
	return NewLineByPoint(NewPoint(x1, y1), NewPoint(x2, y2))
}

func NewRect(left float64, top float64, right float64, bottom float64) Rect {
	return Rect{Left: left, Top: top, Right: right, Bottom: bottom}
}

func NewRectByPoint(ptA Point, ptB Point) Rect {
	left := ptA.X
	top := ptA.Y
	right := ptB.X
	bottom := ptB.Y
	return NewRect(left, top, right, bottom)
}

func (line Line) Rect() Rect {
	return NewRectByPoint(line.PointA, line.PointB)
}

//Exclusion experiment to see whether the line intersects the diagonal of a rectangle
func (line Line) IsRectCross(cl Line) bool {
	ret := math.Min(line.PointA.X, line.PointB.X) <= math.Max(cl.PointA.X, cl.PointB.X) &&
		math.Min(cl.PointA.X, cl.PointB.X) <= math.Max(line.PointA.X, line.PointB.X) &&
		math.Min(line.PointA.Y, line.PointB.Y) <= math.Max(cl.PointA.Y, cl.PointB.Y) &&
		math.Min(cl.PointA.Y, cl.PointB.Y) <= math.Max(line.PointA.Y, line.PointB.Y)

	return ret
}

//Straddles the judge, whether the intersection
func (line Line) IsLineSegmentCross(cl Line) bool {
	//Same endpoint
	if line.PointA.X == cl.PointA.X && line.PointA.Y == cl.PointA.Y ||
		line.PointA.X == cl.PointB.X && line.PointA.Y == cl.PointB.Y ||
		line.PointB.X == cl.PointA.X && line.PointB.Y == cl.PointA.Y ||
		line.PointB.X == cl.PointB.X && line.PointB.Y == cl.PointB.Y {
		return true
	}

	f1 := line.PointA.X*(cl.PointA.Y-line.PointB.Y) +
		line.PointB.X*(line.PointA.Y-cl.PointA.Y) +
		cl.PointA.X*(line.PointB.Y-line.PointA.Y)

	f2 := line.PointA.X*(cl.PointB.Y-line.PointB.Y) +
		line.PointB.X*(line.PointA.Y-cl.PointB.Y) +
		cl.PointB.X*(line.PointB.Y-line.PointA.Y)

	line1 := int64(f1)
	line2 := int64(f2)
	if ((line1 ^ line2) >= 0) && !(line1 == 0 && line2 == 0) {
		return false
	}

	f1 = cl.PointA.X*(line.PointA.Y-cl.PointB.Y) +
		cl.PointB.X*(cl.PointA.Y-line.PointA.Y) +
		line.PointA.X*(cl.PointB.Y-cl.PointA.Y)

	f2 = cl.PointA.X*(line.PointB.Y-cl.PointB.Y) +
		cl.PointB.X*(cl.PointA.Y-line.PointB.Y) +
		line.PointB.X*(cl.PointB.Y-cl.PointA.Y)

	line1 = int64(f1)
	line2 = int64(f2)
	if ((line1 ^ line2) >= 0) && !(line1 == 0 && line2 == 0) {
		return false
	}

	return true
}

//Solving the intersection of two segments
func (line Line) GetCrossPoint(cl Line) (Point, bool) {
	if line.IsRectCross(cl) {
		if line.IsLineSegmentCross(cl) {
			//For the intersection
			tmpLeft := (cl.PointB.X-cl.PointA.X)*(line.PointA.Y-line.PointB.Y) -
				(line.PointB.X-line.PointA.X)*(cl.PointA.Y-cl.PointB.Y)
			tmpRight := (line.PointA.Y-cl.PointA.Y)*(line.PointB.X-line.PointA.X)*
				(cl.PointB.X-cl.PointA.X) + cl.PointA.X*(cl.PointB.Y-cl.PointA.Y)*
				(line.PointB.X-line.PointA.X) - line.PointA.X*
				(line.PointB.Y-line.PointA.Y)*(cl.PointB.X-cl.PointA.X)
			x := tmpRight / tmpLeft

			tmpLeft = (line.PointA.X-line.PointB.X)*(cl.PointB.Y-cl.PointA.Y) -
				(line.PointB.Y-line.PointA.Y)*(cl.PointA.X-cl.PointB.X)
			tmpRight = line.PointB.Y*(line.PointA.X-line.PointB.X)*
				(cl.PointB.Y-cl.PointA.Y) + (cl.PointB.X-line.PointB.X)*
				(cl.PointB.Y-cl.PointA.Y)*(line.PointA.Y-line.PointB.Y) -
				cl.PointB.Y*(cl.PointA.X-cl.PointB.X)*(line.PointB.Y-line.PointA.Y)
			y := tmpRight / tmpLeft

			return NewPoint(x, y), true
		}
	}

	return NewPoint(0, 0), false
}

//Determine whether two rectangles intersect
func (rect Rect) IsCross(crect Rect) bool {
	line := NewLine(rect.Left, rect.Top, rect.Right, rect.Bottom)
	cl := NewLine(crect.Left, crect.Top, crect.Right, crect.Bottom)
	return line.IsRectCross(cl)
}
