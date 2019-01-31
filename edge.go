package intersect

import "math"

// Edge is used for simplyfiying the collision code, as lines, segments, and rays all have very similar code
type Edge interface {
	ToLine() Line
	IsPointOn(Vector) bool
}

// used for edge intersection
func (l1 Line) intersectLine(l2 Line) (Vector, bool) {
	if l1.IsVertical() != l2.IsVertical() {
		if l1.IsVertical() {
			return Vector{l1.xVal, l2.EvalY(l1.xVal)}, true
		}
		return Vector{l2.xVal, l1.EvalY(l2.xVal)}, true
	}
	x := (l2.yInt - l1.yInt) / (l1.slope - l2.slope)
	return Vector{x, l1.EvalY(x)}, !l1.IsParallel(l2)
}

// Intersect returns the intersection between two line segments, and a boolean if it collides
func Intersect(e1, e2 Edge) (Vector, bool) {
	l1 := e1.ToLine()
	l2 := e2.ToLine()
	p, didLinesIntersect := l1.intersectLine(l2)
	if didLinesIntersect {
		if e1.IsPointOn(p) && e2.IsPointOn(p) {
			return p, true
		}
	}
	return Vector{math.NaN(), math.NaN()}, false
}
