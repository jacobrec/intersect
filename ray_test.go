package intersect

import (
	"math"
	"testing"
)

func TestNewRay(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(3, 3))
		if r1.slope != 1 {
			t.Error("Failed Test", r1)
		}
		if !r1.point.Equals(NewVector(1, 1)) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(-1, 1))
		if r1.slope != 0 {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(1, 2))
		if !math.IsInf(r1.slope, 1) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("4", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(1, -1))
		if !math.IsInf(r1.slope, -1) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayToLine(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		if !r1.ToLine().Equals(NewLine(NewVector(3, 3), NewVector(4, 4))) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayIsSameOrigin(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(1, 1), NewVector(3, 2))
		if !r1.IsSameOrigin(r2) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(0, 1), NewVector(3, 2))
		if r1.IsSameOrigin(r2) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayEquals(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(1, 1), NewVector(2, 2))
		if !r1.Equals(r2) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(1, 2), NewVector(2, 2))
		if r1.Equals(r2) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayIsPointOn(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		p := NewVector(3, 3)
		if !r1.IsPointOn(p) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		p := NewVector(0, 0)
		if r1.IsPointOn(p) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRay(NewVector(0, 0), NewVector(-1, 0))
		p := NewVector(-2, 0)
		if !r1.IsPointOn(p) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("4", func(t *testing.T) {
		r1 := NewRay(NewVector(0, 0), NewVector(-1, 0))
		p := NewVector(2, 0)
		if r1.IsPointOn(p) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("5", func(t *testing.T) {
		r1 := NewRay(NewVector(0, 0), NewVector(4, -2))
		p := NewVector(6, -3)
		if !r1.IsPointOn(p) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("6", func(t *testing.T) {
		r1 := NewRay(NewVector(0, -9), NewVector(9, 0))
		p := NewVector(6, -3)
		if !r1.IsPointOn(p) {
			t.Error("Failed Test", r1)
		}
	})

}

func TestRayIsIntersectRay(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(1, 1), NewVector(2, 2))
		if r1.IsIntersectRay(r2) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(0, 0), NewVector(0, 2))
		if r1.IsIntersectRay(r2) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(5, 0), NewVector(5, 2))
		if !r1.IsIntersectRay(r2) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayIsIntersectLine(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		l1 := NewLine(NewVector(1, 1), NewVector(2, 2))
		if r1.IsIntersectLine(l1) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(1, 2))
		l1 := NewLine(NewVector(1, 0), NewVector(2, 0))
		if r1.IsIntersectLine(l1) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("3", func(t *testing.T) {
		r1 := NewRay(NewVector(0, -1), NewVector(2, 2))
		l1 := NewLine(NewVector(1, 1), NewVector(2, 2))
		if !r1.IsIntersectLine(l1) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayIntersectionLine(t *testing.T){
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(0, 0), NewVector(2, 2))
		l1 := NewLine(NewVector(0, 5), NewVector(1, 5))
		if !r1.IntersectionLine(l1).Equals(NewVector(5, 5)) {
			t.Error("Failed Test", r1)
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(2, 1), NewVector(2, 2))
		l1 := NewLine(NewVector(0, 5), NewVector(1, 5))
		if !r1.IntersectionLine(l1).Equals(NewVector(2, 5)) {
			t.Error("Failed Test", r1)
		}
	})
}

func TestRayIntersectionRay(t *testing.T){
	t.Run("1", func(t *testing.T) {
		r1 := NewRay(NewVector(0, 0), NewVector(4, -2))
		r2 := NewRay(NewVector(0, -9), NewVector(9, 0))
		if !r1.IntersectionRay(r2).Equals(NewVector(6, -3)) {
			t.Error("Failed Test", r1, r2, r1.IntersectionRay(r2))
		}
	})
	t.Run("2", func(t *testing.T) {
		r1 := NewRay(NewVector(1, 1), NewVector(2, 2))
		r2 := NewRay(NewVector(0, 5), NewVector(2, 6))
		if !r1.IntersectionRay(r2).Equals(NewVector(10, 10)) {
			t.Error("Failed Test", r1)
		}
	})
}
