package intersect

import (
	"math"
	"testing"
)

func TestEquals(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		if !NewVectorMagDir(10, math.Pi).Equals(NewVectorMagDir(10, math.Pi)) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		nz := math.Copysign(0, -1)
		if !NewVector(0, 0).Equals(NewVector(nz, nz)) {
			t.Error("Failed Test")
		}
	})
}

func TestNewVectorMagDir(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := NewVectorMagDir(10, math.Pi)
		e := NewVector(-10, 0)
		if !v.Equals(e) {
			t.Error("Failed Test: ", v, e)
			t.Error("inner:", floatEquals(v.X, e.X), floatEquals(v.Y, e.Y))
		}
	})
	t.Run("2", func(t *testing.T) {
		v := NewVectorMagDir(10, math.Pi/6)
		e := NewVector(5*math.Sqrt(3), 5)
		if !v.Equals(e) {
			t.Error("Failed Test: ", v, e)
		}
	})
}

func TestNewVector(t *testing.T) {
	v := NewVector(49, 7)
	if v.X != 49 || v.Y != 7 {
		t.Error("Failed Test")
	}
}

func TestDot(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		v1 := Vector{-4, -9}
		v2 := Vector{-1, 2}
		if v1.Dot(v2) != -14 {
			t.Error("Failed Test")
		}
	})

}

func TestLen2(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v1 := Vector{0, -9}
		if v1.Len2() != 81 {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v1 := Vector{5, 0}
		if v1.Len2() != 25 {
			t.Error("Failed Test")
		}
	})
}

func TestLen(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v1 := Vector{-4, -9}
		if v1.Len() != math.Sqrt(97) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v1 := Vector{3, 4}
		if v1.Len() != 5 {
			t.Error("Failed Test")
		}
	})
}

func TestAdd(t *testing.T) {
	v1 := NewVector(3, 8)
	v2 := NewVector(-2, 8)
	v3 := NewVector(2, -9)
	t.Run("1", func(t *testing.T) {
		v := v1.Add(v2)
		if v != NewVector(1, 16) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v := v2.Add(v3)
		if v != NewVector(0, -1) {
			t.Error("Failed Test")
		}
	})
	t.Run("3", func(t *testing.T) {
		v := v3.Add(v1)
		if v != NewVector(5, -1) {
			t.Error("Failed Test")
		}
	})

}

func TestSub(t *testing.T) {
	v1 := NewVector(3, 8)
	v2 := NewVector(-2, 8)
	v3 := NewVector(2, -9)
	t.Run("1", func(t *testing.T) {
		v := v1.Sub(v2)
		if v != NewVector(5, 0) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v := v2.Sub(v3)
		if v != NewVector(-4, 17) {
			t.Error("Failed Test")
		}
	})
	t.Run("3", func(t *testing.T) {
		v := v3.Sub(v1)
		if v != NewVector(-1, -17) {
			t.Error("Failed Test")
		}
	})
}

func TestNorm(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := NewVector(5, 5)
		if !v.Norm().Equals(NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2)) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v := NewVector(3, 4)
		if !v.Norm().Equals(NewVector(0.6, 0.8)) {
			t.Error("Failed Test")
		}
	})
}

func TestLim(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := NewVector(3, 4)
		if !v.Lim(2).Equals(NewVector(1.2, 1.6)) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v := NewVector(67, 4)
		if !v.Lim(1).Equals(v.Norm()) {
			t.Error("Failed Test")
		}
	})
	t.Run("3", func(t *testing.T) {
		v := NewVector(67, 4)
		if !v.Lim(1).Equals(v.Norm()) {
			t.Error("Failed Test")
		}
	})
}
func TestRotate(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := NewVectorMagDir(1, 0).Rotate(math.Pi)
		if math.Abs(v.Angle()-math.Pi) > 1e-9 {
			t.Error("Failed Test: ", v)
		}
	})
	t.Run("2", func(t *testing.T) {
		v := NewVectorMagDir(1, 0).Rotate(-math.Pi)
		if math.Abs(v.Angle()-math.Pi) > 1e-9 {
			t.Error("Failed Test: ", v)
		}
	})
	t.Run("3", func(t *testing.T) {
		v := NewVectorMagDir(3, 3).Rotate(1)
		if math.Abs(v.Angle()-4) > 1e-9 {
			t.Error("Failed Test: ", v)
		}
	})
	t.Run("4", func(t *testing.T) {
		v := NewVector(3, 4)
		v2 := v.Rotate(math.Pi * 2)
		if !v.Equals(v2) {
			t.Error("Failed Test: ", v, v2)
		}
	})
}

func TestAngle(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := NewVectorMagDir(3, 7)
		if math.Abs(v.Angle()-math.Mod(7, 2*math.Pi)) > 1e-9 {
			t.Error("Failed Test: ", v.Angle())
		}
	})
	t.Run("2", func(t *testing.T) {
		v := NewVectorMagDir(3, 1)
		if math.Abs(v.Angle()-1) > 1e-9 {
			t.Error("Failed Test: ", v.Angle())
		}
	})
	t.Run("3", func(t *testing.T) {
		v := NewVectorMagDir(3, math.Pi)
		if math.Abs(v.Angle()-math.Pi) > 1e-9 {
			t.Error("Failed Test: ", v.Angle())
		}
	})
	t.Run("4", func(t *testing.T) {
		v := NewVectorMagDir(3, 3*math.Pi/4)
		if math.Abs(v.Angle()-3*math.Pi/4) > 1e-9 {
			t.Error("Failed Test: ", v.Angle())
		}
	})
}

func TestScale(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := NewVector(3, 4)
		if !v.Scale(2).Equals(NewVector(6, 8)) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v := NewVector(1, 2)
		if !v.Scale(-4).Equals(NewVector(-4, -8)) {
			t.Error("Failed Test")
		}
	})
}

func TestCollinear(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v1 := NewVector(1, 1)
		v2 := NewVector(2, 2)
		v3 := NewVector(3, 3)
		if !Collinear(v1, v2, v3) {
			t.Error("Failed Test")
		}
	})
	t.Run("2", func(t *testing.T) {
		v1 := NewVector(1, 1)
		v2 := NewVector(2, 2)
		v3 := NewVector(2, 4)
		if Collinear(v1, v2, v3) {
			t.Error("Failed Test")
		}
	})
}
