package GoPolygons

import (
	"math"
)

type Polygon struct {
	Points      []Point
	ContainRect Rect
}

func NewPolygon(pts []Point) *Polygon {
	left := pts[0].X
	right := pts[0].X
	top := pts[0].Y
	bottom := pts[0].Y

	for i := 0; i < len((pts)); i++ {
		left = math.Min(left, pts[i].X)
		right = math.Max(right, pts[i].X)
		top = math.Min(top, pts[i].Y)
		bottom = math.Max(bottom, pts[i].Y)
	}
	return &Polygon{Points: pts, ContainRect: NewRect(left, top, right, bottom)}
}

//Ray method to determine whether the point inside the polygon
func (pgn *Polygon) ptInPolygonRayCasting(pt Point) bool {
	px := pt.X
	py := pt.Y
	flag := false
	ptCount := len(pgn.Points)
	i := 0
	j := ptCount - 1
	for i < ptCount {
		sx := pgn.Points[i].X
		sy := pgn.Points[i].Y
		tx := pgn.Points[j].X
		ty := pgn.Points[j].Y

		// Point coincides with the polygon vertices
		if (sx == px && sy == py) || (tx == px && ty == py) {
			return true
		}

		// Determine whether the two end points of the line on both sides of the ray
		if (sy < py && ty >= py) || (sy >= py && ty < py) {
			// X-ray coordinates on the line with the same Y coordinate of the point
			x := sx + (py-sy)*(tx-sx)/(ty-sy)

			// Point in polygon edge
			if x == px {
				return true
			}

			// Rays pass through the polygon boundary
			if x > px {
				flag = !flag
			}
		}
		j = i
		i++
	}

	// When the number of rays passing through the polygon boundary is odd, the point in the polygon
	return flag
}

func (pgn *Polygon) PtInPolygon(pt Point) bool {
	return pgn.ptInPolygonRayCasting(pt)
}

func (pgn *Polygon) ContainPolygon(cpgn Polygon) (ret bool) {
	if pgn.ContainRect.IsCross(cpgn.ContainRect) { //Envelope box must intersect
		ret = true
		PointCount := len(cpgn.Points)
		for i := 0; i < PointCount; i++ {
			if !pgn.PtInPolygon(cpgn.Points[i]) {
				ret = false
				break
			}
		}
	}

	return ret
}

//Polygons intersect, false disjoint, true intersection
func (pgn *Polygon) IntersectPolygon(cpgn Polygon) bool {
	if !pgn.ContainRect.IsCross(cpgn.ContainRect) { //Envelope box must intersect
		return false
	}

	//Positive points included
	for i := 0; i < len(cpgn.Points); i++ {
		if pgn.PtInPolygon(cpgn.Points[i]) {
			return true
		}
	}

	//Anti-contained point
	for i := 0; i < len(pgn.Points); i++ {
		if cpgn.PtInPolygon(pgn.Points[i]) {
			return true
		}
	}

	//Determine whether the line intersect
	ptPgnCount := len(pgn.Points)
	ptCPgnCount := len(cpgn.Points)
	i := 0
	j := ptPgnCount - 1
	for i < ptPgnCount {
		sx := pgn.Points[i].X
		sy := pgn.Points[i].Y
		tx := pgn.Points[j].X
		ty := pgn.Points[j].Y
		line := NewLine(sx, sy, tx, ty)

		k := 0
		l := ptCPgnCount - 1
		for k < ptCPgnCount {
			csx := cpgn.Points[k].X
			csy := cpgn.Points[k].Y
			ctx := cpgn.Points[l].X
			cty := cpgn.Points[l].Y
			cline := NewLine(csx, csy, ctx, cty)
			if line.IsLineSegmentCross(cline) {
				return true
			}
			l = k
			k++
		}
		j = i
		i++
	}

	return false
}
