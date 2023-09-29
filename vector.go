package pt

import (
	"fmt"
	"math"
)

/*
These coordinates are restricted based on a 2D representation.
The reference coordinate 0, 0, 0 is at the top left corner of the screen.
X has positive going to the right.
Y has positive going down.
Z functions as depth, and has positive going out of the screen toward the player
*/
type Vec struct {
	X, Y, Z float64
}

// Add two Vecs and return the resulting Vec struct
func (v1 Vec) Add(v2 Vec) Vec {
	return Vec{X: v1.X + v2.X, Y: v1.Y + v2.Y, Z: v1.Z + v2.Z}
}

// Subtract v2 from v1 and return the Vec result
func (v1 Vec) Subtract(v2 Vec) Vec {
	return Vec{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z}
}

func (v Vec) Scale(a float64) Vec {
	return Vec{X: a * v.X, Y: a * v.Y, Z: a * v.Z}
}

func roundFloat(f float64) float64 {
	return math.Round(1000*f) / 1000
}

func (v Vec) ToInt() (int, int, int) {
	return int(math.Round(v.X)), int(math.Round(v.Y)), int(math.Round(v.Z))
}

func (v Vec) ToInt64() (int64, int64, int64) {
	return int64(math.Round(v.X)), int64(math.Round(v.Y)), int64(math.Round(v.Z))
}

func (v Vec) ToString() string {
	return fmt.Sprintf("%.2f,%.2f,%.2f", v.X, v.Y, v.Z)
}

// Round each dimension to the nearest 1000th to avoid arithmetic
// errors with floating points from propagating
func (v Vec) Thousandths() Vec {
	return Vec{X: roundFloat(v.X), Y: roundFloat(v.Y), Z: roundFloat(v.Z)}
}

// Uses a vector rotate trick to rotate the 2D vector of just X and Y
// around 0, 0.
func (v Vec) RotateXY90() Vec {
	return Vec{X: -v.Y, Y: v.X, Z: v.Z}
}

// Uses a vector rotate trick to rotate the 2D vector of just X and Y
// around 0, 0.
func (v Vec) RotateXY180() Vec {
	return Vec{X: -v.X, Y: -v.Y, Z: v.Z}
}

// Uses a vector rotate trick to rotate the 2D vector of just X and Y
// around 0, 0.
func (v Vec) RotateXY270() Vec {
	return Vec{X: v.Y, Y: -v.X, Z: v.Z}
}

func (v Vec) RotateXY90Around(r Vec) Vec {
	return v.Subtract(r).RotateXY90().Add(r)
}

func (v Vec) RotateXY180Around(r Vec) Vec {
	return v.Subtract(r).RotateXY180().Add(r)
}

func (v Vec) RotateXY270Around(r Vec) Vec {
	return v.Subtract(r).RotateXY90().Add(r)
}
