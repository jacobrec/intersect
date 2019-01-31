package intersect

import "testing"

func TestRectContainsPoint(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRect(0, 0, 1, 1)
		p1 := NewVector(0.5, 0.5)
		if !r1.IsPointIn(p1) {
			t.Error("Test failed", r1, p1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRect(1, 1, 1, 1)
		p1 := NewVector(1.5, 1.5)
		if !r1.IsPointIn(p1) {
			t.Error("Test failed", r1, p1)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRect(1, 1, 1, 1)
		p1 := NewVector(0.5, 0.5)
		if r1.IsPointIn(p1) {
			t.Error("Test failed", r1, p1)
		}
	})
}

func TestRectEquals(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRect(0.5, 0.5, 1, 1)
		r2 := NewRect(0.5, 0.5, 1, 1)
		if !r1.Equals(r2) {
			t.Error("Test failed", r1, r2)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRect(0, 0, 1, 1)
		r2 := NewRect(0, 0, 1, 1)
		if !r1.Equals(r2) {
			t.Error("Test failed", r1, r2)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRect(0, 0, 1, 1)
		r2 := NewRect(0.5, 0.5, 1, 1)
		if r1.Equals(r2) {
			t.Error("Test failed", r1, r2)
		}
	})
}

func TestRectersect(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRect(0, 0, 1, 1)
		r2 := NewRect(0.5, 0.5, 1, 1)
		r3, b := r1.Intersect(r2)
		if !r3.Equals(NewRect(0.5, 0.5, 0.5, 0.5)) {
			t.Error("Test failed", b)
		}
		if !b {
			t.Error("Test failed", r1, r2, r3)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRect(0.5, 0.5, 1, 1)
		r2 := NewRect(0, 0, 1, 1)
		r3, b := r1.Intersect(r2)
		if !r3.Equals(NewRect(0.5, 0.5, 0.5, 0.5)) {
			t.Error("Test failed", b)
		}
		if !b {
			t.Error("Test failed", r1, r2, r3)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRect(0, 0, 2, 2)
		r2 := NewRect(0.5, 0.5, 1, 1)
		r3, b := r1.Intersect(r2)
		if !r3.Equals(NewRect(0.5, 0.5, 1, 1)) {
			t.Error("Test failed", b)
		}
		if !b {
			t.Error("Test failed", r1, r2, r3)
		}
	})
	t.Run("4", func(t *testing.T) {
		r1 := NewRect(0, 0, 2, 2)
		r2 := NewRect(3.5, 3.5, 1, 1)
		r3, b := r1.Intersect(r2)
		if b {
			t.Error("Test failed", r1, r2, r3)
		}
	})
}

func TestRectIntersectEdge(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRect(0, 0, 1, 1)
		e1 := NewLine(NewVector(0.2, 0.2), NewVector(1, 1))
		ps, c := r1.IntersectEdge(e1)
		if c != 2 {
			t.Error("Test failed", ps, c)
		}
		if !(ps[0].Equals(NewVector(0, 0)) && ps[1].Equals(NewVector(1, 1)) ||
			ps[0].Equals(NewVector(1, 1)) && ps[1].Equals(NewVector(0, 0))) {
			t.Error("Test failed", ps, ps[0].Equals(NewVector(0, 0)))
		}
	})

	t.Run("2", func(t *testing.T) {
		r1 := NewRect(0, 0, 1, 1)
		e1 := NewSegment(NewVector(0.5, 0.5), NewVector(1, 1))
		ps, c := r1.IntersectEdge(e1)
		if c != 1 {
			t.Error("Test failed", ps, c)
		}
		if !ps[0].Equals(NewVector(1, 1)) {
			t.Error("Test failed", ps)
		}
	})

}
