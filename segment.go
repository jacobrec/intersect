package intersect

// Segment A line segment is the finite part of a line between two points
type Segment struct {
	start Vector
	end   Vector
}

// NewSegment creates a new line segment bounded by p1, and p2
func NewSegment(p1, p2 Vector) Segment {
	return Segment{}
}
