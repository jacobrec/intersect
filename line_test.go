package intersect

import (
	"math"
	"testing"
)

func TestLineNewLine(t *testing.T) {
	// TODO: test this
}

func TestLineEquals(t *testing.T) {
	t.Run("1", func(t *testing.T) { // diagonal equals
		l1 := NewLine(NewVector(math.Abs(1), 1), NewVector(3, 3))
		l2 := NewLine(NewVector(4, 4), NewVector(3, 3))
		if !l1.Equals(l2) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) { // vertical equals
		l1 := NewLine(NewVector(1, 1), NewVector(1, 3))
		l2 := NewLine(NewVector(1, 3), NewVector(1, 4))
		if !l1.Equals(l2) {
			t.Error("Failed Test", l1, l2)
		}
	})
	t.Run("3", func(t *testing.T) { //horizontal not equals
		l1 := NewLine(NewVector(1, 1), NewVector(3, 1))
		l2 := NewLine(NewVector(4, 3), NewVector(3, 3))
		if l1.Equals(l2) {
			t.Error("Failed Test")
		}
	})
}

func TestLineIsVertical(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(3, 1), NewVector(3, 3))
		if !l1.IsVertical() {
			t.Error("Failed Test")
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 0), NewVector(0, 0))
		if !l1.IsVertical() {
			t.Error("Failed Test", l1)
		}
	})

	t.Run("3", func(t *testing.T) {
		l1 := NewLine(NewVector(5, 3), NewVector(3, 3))
		if l1.IsVertical() {
			t.Error("Failed Test")
		}
	})
}

func TestLineIsPointOn(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(3, 3))
		p := NewVector(10, 10)
		if !l1.IsPointOn(p) {
			t.Error("Failed Test")
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(5, 3), NewVector(3, 3))
		p := NewVector(5, 3)
		if !l1.IsPointOn(p) {
			t.Error("Failed Test")
		}
	})

	t.Run("3", func(t *testing.T) {
		l1 := NewLine(NewVector(5, 3), NewVector(3, 3))
		p := NewVector(10, 239)
		if l1.IsPointOn(p) {
			t.Error("Failed Test")
		}
	})
}

func TestLineIsParallel(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(3, 3))
		l2 := NewLine(NewVector(1, 1), NewVector(3, 3))
		if !l1.IsParallel(l2) {
			t.Error("Failed Test")
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(3, 3))
		l2 := NewLine(NewVector(2, 1), NewVector(4, 3))
		if !l1.IsParallel(l2) {
			t.Error("Failed Test")
		}
	})

	t.Run("3", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(3, 3))
		l2 := NewLine(NewVector(1, 2), NewVector(3, 3))
		if l1.IsParallel(l2) {
			t.Error("Failed Test")
		}
	})

	t.Run("4", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 1), NewVector(0, 3))
		l2 := NewLine(NewVector(1, 1), NewVector(1, 3))
		if !l1.IsParallel(l2) {
			t.Error("Failed Test", l1, l2)
		}
	})
}

func TestLineIsIntersect(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(3, 3))
		l2 := NewLine(NewVector(1, 1), NewVector(3, 3))
		if l1.IsIntersect(l2) {
			t.Error("Failed Test", l1, l2)
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 4), NewVector(2, 3))
		l2 := NewLine(NewVector(2, 8), NewVector(0, 3))
		if !l1.IsIntersect(l2) {
			t.Error("Failed Test")
		}
	})

	t.Run("3", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(3, 3))
		l2 := NewLine(NewVector(2, 1), NewVector(4, 3))
		if l1.IsIntersect(l2) {
			t.Error("Failed Test")
		}
	})
}

func TestLineEvalX(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 0), NewVector(2, 2))
		if l1.EvalX(3) != 3 {
			t.Error("Failed Test")
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 0), NewVector(3, 0))
		if !math.IsNaN(l1.EvalX(3)) {
			t.Error("Failed Test", l1)
		}
	})
}

func TestLineEvalY(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 0), NewVector(2, 2))
		if l1.EvalY(3) != 3 {
			t.Error("Failed Test")
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 0), NewVector(3, 0))
		if l1.EvalY(3) != 0 {
			t.Error("Failed Test")
		}
	})
}

func TestLineIntersection(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := NewLine(NewVector(0, 2), NewVector(4, 2))
		l2 := NewLine(NewVector(2, 0), NewVector(2, 4))
		if !l1.Intersection(l2).Equals(NewVector(2, 2)) {
			t.Error("Failed Test")
		}
	})

	t.Run("2", func(t *testing.T) {
		l1 := NewLine(NewVector(1, 1), NewVector(2, 2))
		l2 := NewLine(NewVector(3, 2), NewVector(3, 2))
		if !l1.Intersection(l2).Equals(NewVector(3, 3)) {
			t.Error("Failed Test", l1, l2)
		}
	})
}
