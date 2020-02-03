package main

import "math"

type Camera struct {
	origin, lowerLeftCorner, horizontal, vertical Tuple
}

func getCamera(lookFrom, lookAt, up Tuple, fov, aspect float64) Camera {
	var c Camera
	var u, v, w Tuple
	theta := fov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	c.origin = lookFrom
	w = (lookFrom.Subtract(lookAt)).Normalize()
	u = up.Cross(w).Normalize()
	v = w.Cross(u)
	c.lowerLeftCorner = c.origin.Add(u.MulScalar(-(halfWidth))).Add(v.MulScalar(-(halfHeight))).Add(w)
	c.horizontal = u.MulScalar(2 * halfWidth)
	c.vertical = v.MulScalar(2 * halfHeight)

	return c
}

func (c Camera) getRay(s, t float64) Ray {
	return Ray{c.origin, c.lowerLeftCorner.Add(c.horizontal.MulScalar(s)).Add(c.vertical.MulScalar(t)).Subtract(c.origin)}
}
