package intersect

import "math"

// Vector defines a struct for a 2d vector, or a point
type Vector struct {
	X float64
	Y float64
}

// NewVector creates a new vector
func NewVector(x, y float64) Vector {
	return Vector{x, y}
}

// NewVectorMagDir creates a new vector by specifying the magnitude and direction
func NewVectorMagDir(mag, dir float64) Vector {
	return Vector{mag * math.Cos(dir), mag * math.Sin(dir)}
}

// Equals returns true if the x and y components of the vectors are within 1E-9 of each other
func (v1 Vector) Equals(v2 Vector) bool {
	return floatEquals(v1.X, v2.X) && floatEquals(v1.Y, v2.Y)
}

// Len returns the dot product of the vector with another
func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Len2 returns the length of the vector
func (v1 Vector) Len() float64 {
	return math.Sqrt(v1.Len2())
}

// Len2 returns the length of the vector squared. This is faster then Len, and useful for comparisons
func (v1 Vector) Len2() float64 {
	return v1.X*v1.X + v1.Y*v1.Y
}

// Add returns the addition of one vector with another
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

// Sub returns the subtraction of one vector with another
func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

// Norm return a vector of the same angle, but whose length is no 1
func (v1 Vector) Norm() Vector {
	l := v1.Len()
	return Vector{v1.X / l, v1.Y / l}
}

// Lim return a vector of the same angle, but whose length is no longer then the maximum
func (v1 Vector) Lim(max float64) Vector {
	l := v1.Len()
	if l > max {
		return v1.Norm().Scale(max)
	}
	return v1
}

// Rotate return a vector rotated by the given angle in radians counter clockwise
func (v1 Vector) Rotate(ang float64) Vector {
	return NewVectorMagDir(v1.Len(), v1.Angle()+ang)
}

// Angle return an angle in radians the vector is at
func (v1 Vector) Angle() float64 {
	a := math.Mod(math.Atan2(v1.Y, v1.X), 2*math.Pi)
	if a < 0 {
		a += math.Pi * 2
	}
	return a
}

// Scale return a vector in the same direction, but with the magnitude multiplied by the scaling factor
func (v1 Vector) Scale(sf float64) Vector {
	return Vector{v1.X * sf, v1.Y * sf}
}

// Collinear return true iff a, b, and c all lie on the same line."
func Collinear(a, b, c Vector) bool {
	return (b.X-a.X)*(c.Y-a.Y) == (c.X-a.X)*(b.Y-a.Y)
}

const fEPSILON float64 = 1e-9

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < fEPSILON {
		return true
	}
	return false
}
