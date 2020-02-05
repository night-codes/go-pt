package main

// Ray struct represents a ray with origin and a direction
type Ray struct {
	origin, direction Tuple
}

// Position returns point after traveling distance `t` along a vector
func (ray Ray) Position(t float64) Tuple {
	return ray.origin.Add(ray.direction.MulScalar(t))
}
