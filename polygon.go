package intersect

// Polygon is a generic polygon struct
type Polygon struct {
	Points []Vector
}

// NewPoly creates a new polygon, assumes points are clockwise
func NewPoly(points []Vector) Polygon {
	return Polygon{points}
}

// NewPolyFromSegments creates a new polygon from line segments, assumes segments connect clockwise
func NewPolyFromSegments(ss []Segment) Polygon {
	points := make([]Vector, 0)
	for _, s := range ss {
		points = append(points, s.end)
	}
	return NewPoly(points)
}

// IntersectEdge calculates where a straight edge intersects with the rectangle. It returns 0, 1, or 2 points in an array, as well as a number indicating how many collisions occured
func (p1 Polygon) IntersectEdge(e2 Edge) ([]Vector, int) {
	ps := make([]Vector, 0)
	c := 0
	ss := make([]Segment, 1)
	ss[0] = NewSegment(p1.Points[0], p1.Points[len(p1.Points)-1])
	for i := range p1.Points {
		if i != 0 {
			ss = append(ss, NewSegment(p1.Points[i-1], p1.Points[i]))
		}
	}

	prior := make([]Vector, 0)
	for _, s := range ss {
		p, b := s.Intersect(e2)
		if b && notInPrior(prior, p) {
			ps = append(ps, p)
			prior = append(prior, p)
			c++
		}
	}
	return ps, c
}

func notInPrior(ps []Vector, p Vector) bool {
	for _, v := range ps {
		if p.Equals(v) {
			return false
		}
	}
	return true
}
