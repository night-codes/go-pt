package main

import (
	"math/rand"
)

func RandFloat() float64 {
	return rand.Float64()
}

func RandInUnitSphere() Tuple {
	p := Tuple{0, 0, 0, 0}
	for {
		p = (Tuple{RandFloat(), RandFloat(), RandFloat(), 1}.Subtract(Tuple{0.5, 0.5, 0.5, 0})).MulScalar(2.0)
		if p.Magnitude()*p.Magnitude() < 1.0 {
			break
		}
	}
	return p
}
