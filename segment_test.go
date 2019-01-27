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
