package main

import (
	"math"
	"math/rand"
)

const (
	Metal      = iota
	Lambertian = iota
	Dielectric = iota
	Emission   = iota
	Plastic    = iota
)

type Material struct {
	material  int
	albedo    Color
	roughness float64
	ior       float64
	checkered bool
}

func (m Material) Scatter(r Ray, rec HitRecord, attenuation *Color, scattered *Ray, generator rand.Rand) bool {
	if m.material == Lambertian {
		target := rec.p.Add(rec.normal).Add(RandInUnitSphere(generator))
		*scattered = Ray{rec.p, target.Subtract(rec.p)}
		if m.checkered {
			if (int(math.Floor(rec.p.x/0.23))+int(math.Floor(rec.p.y/0.23))+int(math.Floor(rec.p.z/0.23)))%2 == 0 {
				*attenuation = Color{0.7, 0.7, 0.7}
			} else {
				*attenuation = m.albedo
			}
		} else {
			*attenuation = m.albedo
		}
		return true
	} else if m.material == Metal {
		reflected := (r.direction.Normalize()).Reflection(rec.normal)
		*scattered = Ray{rec.p, reflected.Add(RandInUnitSphere(generator).MulScalar(m.roughness))}
		*attenuation = m.albedo
		return (scattered.direction.Dot(rec.normal) > 0)
	} else if m.material == Dielectric {
		var outwardNormal Tuple
		var refracted Tuple

		var niOverNt float64
		var reflectProbability float64
		var cosine float64

		*attenuation = m.albedo
		reflected := r.direction.Reflection(rec.normal)

		if r.direction.Dot(rec.normal) > 0 {
			outwardNormal = rec.normal.MulScalar(-1)
			niOverNt = m.ior
			cosine = m.ior * r.direction.Dot(rec.normal) / r.direction.Magnitude()
		} else {
			outwardNormal = rec.normal
			niOverNt = 1.0 / m.ior
			cosine = -(r.direction.Dot(rec.normal) / r.direction.Magnitude())
		}

		if r.direction.Refraction(outwardNormal, niOverNt, &refracted) {
			reflectProbability = Schlick(cosine, m.ior)
		} else {
			reflectProbability = 1.0
		}

		if RandFloat(generator) < reflectProbability {
			*scattered = Ray{rec.p, reflected.Add(RandInUnitSphere(generator).MulScalar(m.roughness))}
		} else {
			*scattered = Ray{rec.p, refracted.Add(RandInUnitSphere(generator).MulScalar(m.roughness))}
		}

		return true
	} else if m.material == Emission {
		return true
	} else if m.material == Plastic {
		var outwardNormal Tuple
		var refracted Tuple

		var niOverNt float64
		var reflectProbability float64
		var cosine float64

		*attenuation = m.albedo
		reflected := r.direction.Reflection(rec.normal)

		if r.direction.Dot(rec.normal) > 0 {
			outwardNormal = rec.normal.MulScalar(-1)
			niOverNt = m.ior
			cosine = m.ior * r.direction.Dot(rec.normal) / r.direction.Magnitude()
		} else {
			outwardNormal = rec.normal
			niOverNt = 1.0 / m.ior
			cosine = -(r.direction.Dot(rec.normal) / r.direction.Magnitude())
		}

		if r.direction.Refraction(outwardNormal, niOverNt, &refracted) {
			reflectProbability = Schlick(cosine, m.ior)
		} else {
			reflectProbability = 1.0
		}

		if RandFloat(generator) < reflectProbability {
			*scattered = Ray{rec.p, reflected.Add(RandInUnitSphere(generator).MulScalar(m.roughness))}
		} else {
			target := rec.p.Add(rec.normal).Add(RandInUnitSphere(generator))
			*scattered = Ray{rec.p, target.Subtract(rec.p)}
		}

		return true
	}
	return false
}
