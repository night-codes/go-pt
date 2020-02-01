package main

import "math"

type HitRecord struct {
	t        float64
	p        Tuple
	normal   Tuple
	material Material
}

type HittableList struct {
	hits []Sphere
}

func (h *HittableList) hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	tempRec := HitRecord{0.0, Tuple{0, 0, 0, 0}, Tuple{0, 0, 0, 0}, Material{-1, Color{0, 0, 0}, 0, 0}}
	hitAnything := false
	closestSoFar := tMax
	for i := 0; i < len(h.hits); i++ {
		if h.hits[i].hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}
	return hitAnything
}

func (s *Sphere) hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.origin.Subtract(s.origin)
	a := r.direction.Dot(r.direction)
	b := 2.0 * oc.Dot(r.direction)
	c := oc.Dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c
	if discriminant > 0.0 {
		temp := (-b - math.Sqrt(discriminant)) / (2.0 * a)
		if temp < tMax && temp > tMin {
			*&rec.t = temp
			*&rec.p = r.Position(rec.t)
			*&rec.normal = (rec.p.Subtract(s.origin)).DivScalar(s.radius)
			*&rec.material = s.material
			return true
		}
		temp = (-b + math.Sqrt(discriminant)) / (2.0 * a)
		if temp < tMax && temp > tMin {
			*&rec.t = temp
			*&rec.p = r.Position(rec.t)
			*&rec.normal = (rec.p.Subtract(s.origin)).DivScalar(s.radius)
			*&rec.material = s.material
			return true
		}
	}
	return false
}
