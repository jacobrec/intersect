package intersect

// Ray A ray is infinate in length on one side, it is defined as a starting point, and another point through with to pass
type Ray struct {
	slope float64
	point Vector
}

// NewRay creates a new ray, travelling from p1 in the direction of p2
func NewRay(p1, p2 Vector) Ray {
	return Ray{(p2.Y - p1.Y) / (p2.X - p1.X), p1}
}

// ToLine produces a Line that travels along the ray
func (r1 Ray) ToLine() Line {
	return Line{}
}

// Equals checks if the rays are equal within 1e-9
func (r1 Ray) Equals(r2 Ray) bool {
	return false
}

// IsPointOn returns true if the following point is on the ray
func (r1 Ray) IsPointOn(p1 Vector) bool {
	return false
}

// IsSameOrigin returns true if the following ray starts from the same point
func (r1 Ray) IsSameOrigin(r2 Ray) bool {
	return false
}

// IsIntersectRay returns true if the following ray intersects
func (r1 Ray) IsIntersectRay(r2 Ray) bool {
	return false
}

// IsIntersectLine returns true if the following line intersects
func (r1 Ray) IsIntersectLine(l1 Line) bool {
	return false
}
