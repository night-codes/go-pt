package main

import "math"

type HitRecord struct {
	t        float64
	p        Tuple
	normal   Tuple
	material Material
}

type HittableList struct {
	sphereHits   []Sphere
	triangleHits []Triangle
}

func (h *HittableList) hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax
	for i := 0; i < len(h.sphereHits); i++ {
		if h.sphereHits[i].hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}
	for i := 0; i < len(h.triangleHits); i++ {
		if h.triangleHits[i].hit(r, tMin, closestSoFar, &tempRec) {
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

func (tri *Triangle) hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	vertex0 := tri.position.vertex0
	vertex1 := tri.position.vertex1
	vertex2 := tri.position.vertex2
	edge1 := vertex1.Subtract(vertex0)
	edge2 := vertex2.Subtract(vertex0)
	h := r.direction.Cross(edge2)
	a := edge1.Dot(h)
	if a > -Epsilon && a < Epsilon {
		return false
	}
	f := 1.0 / a
	s := r.origin.Subtract(vertex0)
	u := f * s.Dot(h)
	if u < 0.0 || u > 1.0 {
		return false
	}
	q := s.Cross(edge1)
	v := f * r.direction.Dot(q)
	if v < 0.0 || u+v > 1.0 {
		return false
	}
	t := f * edge2.Dot(q)
	if t < tMax && t > tMin {
		*&rec.p = r.origin.Add(r.direction.MulScalar(t))
		*&rec.t = t
		*&rec.material = tri.material
		// *&rec.normal = Tuple{0, 1, 0, 1}
		// *&rec.normal = edge1.Cross(edge2)
		*&rec.normal = tri.normal
		return true
	}
	return false
}
