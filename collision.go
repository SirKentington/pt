package pt

import "math"

type Hitbox struct {
	Solid bool // Should this hitbox be included in collision detection
	// Two points representing corner vertices of the hitbox in 3D space
	BV1 Vec
	BV2 Vec
}

// Check if two "lines" in a dimension (d1 and d2) overlap.
func dimensionalOverlap(d11, d12, d21, d22 float64) bool {
	if math.Max(d11, d12) < math.Min(d21, d22) {
		return false
	}
	if math.Max(d21, d22) < math.Min(d11, d12) {
		return false
	}
	return true
}

// Return true if two hitboxes are currently colliding
// where colliding means either they overlap or their
// surfaces are touching
func AreColliding(h1, h2 Hitbox) bool {
	return dimensionalOverlap(
		h1.BV1.X, h1.BV2.X, h2.BV1.X, h2.BV2.X,
	) && dimensionalOverlap(
		h1.BV1.Y, h1.BV2.Y, h2.BV1.Y, h2.BV2.Y,
	) && dimensionalOverlap(
		h1.BV1.Z, h1.BV2.Z, h2.BV1.Z, h2.BV2.Z,
	)
}
