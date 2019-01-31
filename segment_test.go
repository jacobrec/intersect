package intersect

import "testing"

func TestSegToLine(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		l1 := NewLine(NewVector(0, 0), NewVector(1, 1))
		if !s1.ToLine().Equals(l1) {
			t.Error("Failed Test", s1)
		}
	})
	t.Run("2", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 1), NewVector(2, 2))
		l1 := NewLine(NewVector(0, 0), NewVector(1, 1))
		if s1.ToLine().Equals(l1) {
			t.Error("Failed Test", s1)
		}
	})
}

func TestSegIsPointOn(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		p1 := NewVector(1, 1)
		if !s1.IsPointOn(p1) {
			t.Error("Failed Test", s1)
		}
	})
	t.Run("2", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		p1 := NewVector(3, 3)
		if s1.IsPointOn(p1) {
			t.Error("Failed Test", s1)
		}
	})
}

func TestSegIntersectSeg(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewSegment(NewVector(1, 0), NewVector(1, 2))
		p, b := s1.Intersect(s2)
		if !b && !p.Equals(NewVector(1, 1)) {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("2", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewSegment(NewVector(3, 0), NewVector(3, 2))
		p, b := s1.Intersect(s2)
		if b {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("3", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewSegment(NewVector(-1, 0), NewVector(-1, 2))
		p, b := s1.Intersect(s2)
		if b {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("4", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewSegment(NewVector(1, 0), NewVector(3, 2))
		p, b := s1.Intersect(s2)
		if b {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
}

func TestSegIntersectRay(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewRay(NewVector(1, 0), NewVector(1, 2))
		p, b := s1.Intersect(s2)
		if !b && !p.Equals(NewVector(1, 1)) {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("2", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(3, 3))
		s2 := NewRay(NewVector(3, 0), NewVector(3, 2))
		p, b := s1.Intersect(s2)
		if !b && !p.Equals(NewVector(3, 3)) {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("3", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewRay(NewVector(-1, 0), NewVector(-1, 2))
		p, b := s1.Intersect(s2)
		if b {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("4", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewRay(NewVector(1, 0), NewVector(3, 2))
		p, b := s1.Intersect(s2)
		if b {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
}

func TestSegIntersectLine(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewLine(NewVector(1, 0), NewVector(1, 2))
		p, b := s1.Intersect(s2)
		if !b && !p.Equals(NewVector(1, 1)) {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("2", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(3, 3))
		s2 := NewLine(NewVector(3, 0), NewVector(3, 2))
		p, b := s1.Intersect(s2)
		if !b && !p.Equals(NewVector(3, 3)) {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("3", func(t *testing.T) {
		s1 := NewSegment(NewVector(-1, -1), NewVector(2, 2))
		s2 := NewLine(NewVector(-1, 0), NewVector(-1, 2))
		p, b := s1.Intersect(s2)
		if !b && !p.Equals(NewVector(-1, -1)) {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
	t.Run("4", func(t *testing.T) {
		s1 := NewSegment(NewVector(0, 0), NewVector(2, 2))
		s2 := NewLine(NewVector(1, 0), NewVector(3, 2))
		p, b := s1.Intersect(s2)
		if b {
			t.Error("Failed Test", s1, s2, p, b)
		}
	})
}
