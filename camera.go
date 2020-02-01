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
	origin := lookFrom
	w = (lookFrom.Subtract(lookAt)).Normalize()
	u = up.Cross(w).Normalize()
	v = w.Cross(u)
	c.lowerLeftCorner = origin.Subtract(u.MulScalar(halfWidth)).Subtract(v.MulScalar(halfHeight).Subtract(w))
	c.horizontal = u.MulScalar(halfWidth).MulScalar(2)
	c.vertical = v.MulScalar(halfHeight).MulScalar(2)

	return c
}

func (c Camera) getRay(s, t float64) Ray {
	return Ray{c.origin, c.lowerLeftCorner.Add(c.horizontal.MulScalar(s).Add(c.vertical.MulScalar(t).Subtract(c.origin)))}
}
