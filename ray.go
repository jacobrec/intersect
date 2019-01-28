package intersect

import "math"

// Ray A ray is infinate in length on one side, it is defined as a starting point, and another point through with to pass
type Ray struct {
	slope         float64
	point         Vector
	isXIncreasing bool
}

// NewRay creates a new ray, travelling from p1 in the direction of p2
func NewRay(p1, p2 Vector) Ray {
	return Ray{(p2.Y - p1.Y) / (p2.X - p1.X), p1, p2.X > p1.X}
}

// ToLine produces a Line that travels along the ray
func (r1 Ray) ToLine() Line {
	if math.IsInf(r1.slope, 1) {
		return NewLine(r1.point, NewVector(r1.point.X, r1.point.Y+1))
	}
	return NewLine(r1.point, NewVector(r1.point.X+1, r1.point.Y+r1.slope))
}

// Equals checks if the rays are equal within 1e-9
func (r1 Ray) Equals(r2 Ray) bool {
	return floatEquals(r1.slope, r2.slope) && r1.point.Equals(r2.point) && r1.isXIncreasing == r2.isXIncreasing
}

// IsPointOn returns true if the following point is on the ray
func (r1 Ray) IsPointOn(p1 Vector) bool {
	if math.IsInf(r1.slope, 1) {
		return floatEquals(p1.X, r1.point.X) && math.Copysign(1, (p1.Y-r1.point.Y)) == math.Copysign(1, r1.slope)
	}
	return r1.ToLine().IsPointOn(p1) && (p1.X > r1.point.X == r1.isXIncreasing)
}

// IsSameOrigin returns true if the following ray starts from the same point
func (r1 Ray) IsSameOrigin(r2 Ray) bool {
	return r1.point.Equals(r2.point)
}

// IntersectRay returns the point where the two rays intersect, if they dont intersect, the result is undefined, and the bool is false
func (r1 Ray) IntersectRay(r2 Ray) (Vector, bool) {
	l1 := r1.ToLine()
	l2 := r2.ToLine()
	p, didLineIntersect := l1.Intersect(l2)
	if didLineIntersect {
		if r1.IsPointOn(p) && r2.IsPointOn(p) {
			return p, true
		}
	}
	return Vector{math.NaN(), math.NaN()}, false
}

// IntersectLine returns the point where the ray and line intersect, if they dont intersect, the result is undefined, and the bool is false
func (r1 Ray) IntersectLine(l1 Line) (Vector, bool) {
	l2 := r1.ToLine()
	p, didLineIntersect := l1.Intersect(l2)
	if didLineIntersect {
		if r1.IsPointOn(p) {
			return p, true
		}
	}
	return Vector{math.NaN(), math.NaN()}, false
}

// IntersectSegment returns the point the segment and ray intersect, and a boolean which determines if they intersect
func (r1 Ray) IntersectSegment(s2 Segment) (Vector, bool) {
	return s2.IntersectRay(r1)
}
