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

// Intersect returns the point the segment and edge intersect, and a boolean which determines if they intersect
func (s1 Segment) Intersect(e2 Edge) (Vector, bool) {
	return Intersect(s1, e2)
}
