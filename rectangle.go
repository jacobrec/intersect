package intersect

import "math"

// Rectangle is a simple geometry struct
type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// NewRect returns a new rectangle
func NewRect(x, y, w, h float64) Rectangle {
	return Rectangle{x, y, w, h}
}

// IsPointIn returns true if the point is in the rectangle
func (r1 Rectangle) IsPointIn(p2 Vector) bool {
	return within(r1.X, p2.X, r1.X+r1.Width) && within(r1.Y, p2.Y, r1.Y+r1.Height)
}

// Equals returns true if all the parameters are close enough to the other rectangle
func (r1 Rectangle) Equals(r2 Rectangle) bool {
	return floatEquals(r1.X, r2.X) && floatEquals(r1.Y, r2.Y) && floatEquals(r1.Width, r2.Width) && floatEquals(r1.Height, r2.Height)
}

// This function assumes x1 < x2
func overlapCollinearLines(x1, w1, x2, w2 float64) (float64, float64) {
	var w float64
	if x2 > x1+w1 { // no overlap
		w = 0
	} else if x1+w1 > x2+w2 { // line 1 encompasses line 2
		w = w2
	} else { // normal case
		w = x1 + w1 - x2
	}
	return x2, w
}
func intersectNormalizedRectangles(r1, r2 Rectangle) Rectangle {
	x, w := overlapCollinearLines(r1.X, r1.Width, r2.X, r2.Width)
	y, h := overlapCollinearLines(r1.Y, r1.Height, r2.Y, r2.Height)
	return NewRect(x, y, w, h)
}

// Intersect returns rectangle formed by the intersection as well as a boolean indicating if there was an itersection
func (r1 Rectangle) Intersect(r2 Rectangle) (Rectangle, bool) {
	var x1, y1, w1, h1 float64
	var x2, y2, w2, h2 float64

	if r1.X < r2.X {
		x1 = r1.X
		w1 = r1.Width
		x2 = r2.X
		w2 = r2.Width
	} else {
		x2 = r1.X
		w2 = r1.Width
		x1 = r2.X
		w1 = r2.Width
	}

	if r1.Y < r2.Y {
		y1 = r1.Y
		h1 = r1.Height
		y2 = r2.Y
		h2 = r2.Height
	} else {
		y2 = r1.Y
		h2 = r1.Height
		y1 = r2.Y
		h1 = r2.Height
	}
	r := intersectNormalizedRectangles(NewRect(x1, y1, w1, h1), NewRect(x2, y2, w2, h2))
	return r, r.Width > 0 && r.Height > 0
}

// IntersectEdge calculates where a straight edge intersects with the rectangle. It returns 0, 1, or 2 points in an array, as well as a number indicating how many collisions occured
func (r1 Rectangle) IntersectEdge(e2 Edge) ([2]Vector, int) {
	return [2]Vector{NewVector(math.NaN(), math.NaN()), NewVector(math.NaN(), math.NaN())}, 0
}
