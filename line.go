package intersect

import "math"

// Line A line is infinate in length, and is defined by any two points
type Line struct {
	slope float64
	yInt  float64
	xVal  float64 // only used if slope is inf
}

// NewLine creates a new line passing through p1 and p2, if p1 == p2, then a vertical line will be created
func NewLine(p1, p2 Vector) Line {
	xVal := p1.X
	slope := (p2.Y - p1.Y) / (p2.X - p1.X)
	if math.IsNaN(slope) {
		slope = math.Inf(1)
	}
	yInt := p1.Y - slope*p1.X
	return Line{slope, yInt, xVal}
}

// Equals checks if the two lines are the same, within 1e-9
func (l1 Line) Equals(l2 Line) bool {
	if l1.IsVertical() && l2.IsVertical() {
		return floatEquals(l1.xVal, l2.xVal)
	}
	return floatEquals(l1.slope, l2.slope) && floatEquals(l1.yInt, l2.yInt)
}

// IsPointOn returns true if the point is on the line
func (l1 Line) IsPointOn(p1 Vector) bool {
	if l1.IsVertical() {
		return floatEquals(p1.X, l1.xVal)
	}
	return floatEquals(l1.yInt+l1.slope*p1.X, p1.Y)
}

// IsParallel returns true if the slopes are equal. If the lines are the same line, it will still return true
func (l1 Line) IsParallel(l2 Line) bool {
	return floatEquals(l1.slope, l2.slope) || (l1.IsVertical() && l2.IsVertical())
}

// IsIntersect returns true, if the two lines intersect, is false if they are the same line
func (l1 Line) IsIntersect(l2 Line) bool {
	return !l1.IsParallel(l2)
}

// IsVertical returns true if the slope of the line is infinate
func (l1 Line) IsVertical() bool {
	return math.IsInf(l1.slope, 0)
}

// EvalX returns the x for a given y on the line
func (l1 Line) EvalX(y float64) float64 {
	if l1.slope == 0 {
		return math.NaN()
	}
	return (y - l1.yInt) / l1.slope
}

// EvalY returns the y for a given x on the line
func (l1 Line) EvalY(x float64) float64 {
	if l1.IsVertical() {
		return math.NaN()
	}
	return l1.yInt + l1.slope*x
}

// Intersect returns the point the lines intersect, result is undefined if there is no intersection
func (l1 Line) Intersect(l2 Line) (Vector, bool) {
	if l1.IsVertical() != l2.IsVertical() {
		if l1.IsVertical() {
			return Vector{l1.xVal, l2.EvalY(l1.xVal)}, true
		}
		return Vector{l2.xVal, l1.EvalY(l2.xVal)}, true
	}
	x := (l2.yInt - l1.yInt) / (l1.slope - l2.slope)
	return Vector{x, l1.EvalY(x)}, !l1.IsParallel(l2)
}
