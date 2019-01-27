package intersect

// Segment A line segment is the finite part of a line between two points
type Segment struct {
	start Vector
	end   Vector
}

// NewSegment creates a new line segment bounded by p1, and p2
func NewSegment(p1, p2 Vector) Segment {
	return Segment{p1, p2}
}

// ToLine creates a new line from this segment
func (s1 Segment) ToLine() Line {
	return NewLine(s1.start, s1.end)
}

// IsPointOn return true iff point c intersects the line segment from a to b.
func (s1 Segment) IsPointOn(p1 Vector) bool {
	if s1.start.X != s1.end.X {
		return within(s1.start.X, p1.X, s1.end.X) && Collinear(s1.start, s1.end, p1)
	}
	return within(s1.start.Y, p1.Y, s1.end.Y) && Collinear(s1.start, s1.end, p1)
}

// within return true iff q is between p and r (inclusive).
func within(p, q, r float64) bool {
	return (p <= q) && (q <= r) || (r <= q) && (q <= p)
}

// Intersect returns the intersection between two line segments, and a boolean if it collides
func (s1 Segment) Intersect(s2 Segment) (Vector, bool) {
	return Vector{}, false
}

// IntersectRay returns the intersection with a ray, and a boolean if it collides
func (s1 Segment) IntersectRay(r2 Ray) (Vector, bool) {
	return Vector{}, false
}

// IntersectLine returns the intersection with a line, and a boolean if it collides
func (s1 Segment) IntersectLine(l2 Line) (Vector, bool) {
	return Vector{}, false
}
