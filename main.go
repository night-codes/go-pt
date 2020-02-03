package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	hsize   = 2000
	vsize   = 1280
	samples = 2048
	depth   = 8
)

func color(r Ray, world *HittableList, d int) Color {
	rec := HitRecord{}
	if world.hit(r, Epsilon, math.MaxFloat64, &rec) {
		var attenuation Color
		var scattered Ray
		if d < depth && rec.material.Scatter(r, rec, &attenuation, &scattered) {
			if rec.material.material == Emission {
				return rec.material.albedo
			} else {
				return attenuation.Mul(color(scattered, world, d+1))
			}
		} else {
			return Color{0, 0, 0}
		}
	} else {
		unit_direction := r.direction.Normalize()
		t := 0.5 * (unit_direction.y + 1.0)
		return Color{1.0, 1.0, 1.0}.MulScalar(1.0 - t).Add(Color{0.5, 0.7, 1.0}.MulScalar(t))
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	canvas := make([]Color, vsize*hsize)
	fmt.Printf("Rendering %dx%d at %d samples\n", hsize, vsize, samples)

	// camera := getCamera(Tuple{0, 0, -1, 0}, Tuple{0, 0, 1, 0}, Tuple{0, 1, 0, 0}, 90, float64(hsize)/float64(vsize))
	cameraPosition := Tuple{0, 0.25, -0.5, 0}
	cameraDirection := Tuple{0, 0.25, -1, 0}
	// cameraPosition := Tuple{-2, 2, 1, 0}
	camera := getCamera(cameraPosition, cameraDirection, Tuple{0, 1, 0, 0}, 75, float64(hsize)/float64(vsize))
	// camera := getCamera(Tuple{0, 0, 0, 0}, Tuple{0, 0, 1, 0}, Tuple{0, -1, 0, 0}, 100, float64(hsize)/float64(vsize))

	fmt.Printf("%v\n", camera)

	list := []Sphere{}
	// list = append(list, Sphere{
	// 	Tuple{0.25, -0.25, 1, 0},
	// 	0.25,
	// 	Material{Dielectric, Color{1, 1, 1}, 0, 1.45, false},
	// })

	// list = append(list, Sphere{
	// 	Tuple{-0.25, -0.25, 1, 0},
	// 	0.25,
	// 	Material{Lambertian, Color{1, 1, 1}, 0, 0, false},
	// })

	// list = append(list, Sphere{
	// 	Tuple{0.25, 0.25, 1, 0},
	// 	0.25,
	// 	Material{Metal, Color{1, 1, 1}, 0, 0, false},
	// })

	// list = append(list, Sphere{
	// 	Tuple{-0.25, 0.25, 1, 0},
	// 	0.25,
	// 	Material{Plastic, Color{1, 1, 1}, 0, 1.45, false},
	// })

	// diel

	list = append(list, Sphere{
		Tuple{0.75, -0.25, 1, 0},
		0.25,
		Material{Dielectric, Color{1, 1, 1}, 0, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{0.25, -0.25, 1, 0},
		0.25,
		Material{Dielectric, Color{1, 1, 1}, 0.02, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{-0.25, -0.25, 1, 0},
		0.25,
		Material{Dielectric, Color{1, 1, 1}, 0.1, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{-0.75, -0.25, 1, 0},
		0.25,
		Material{Dielectric, Color{1, 1, 1}, 0.3, 1.45, false},
	})

	// metal

	list = append(list, Sphere{
		Tuple{0.75, 0.25, 1, 0},
		0.25,
		Material{Metal, Color{1, 1, 1}, 0, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{0.25, 0.25, 1, 0},
		0.25,
		Material{Metal, Color{1, 1, 1}, 0.1, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{-0.25, 0.25, 1, 0},
		0.25,
		Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{-0.75, 0.25, 1, 0},
		0.25,
		Material{Metal, Color{1, 1, 1}, 0.8, 1.45, false},
	})

	// plastic

	list = append(list, Sphere{
		Tuple{0.75, 0.75, 1, 0},
		0.25,
		Material{Plastic, Color{1, 1, 1}, 0, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{0.25, 0.75, 1, 0},
		0.25,
		Material{Plastic, Color{1, 1, 1}, 0.02, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{-0.25, 0.75, 1, 0},
		0.25,
		Material{Plastic, Color{1, 1, 1}, 0.05, 1.45, false},
	})

	list = append(list, Sphere{
		Tuple{-0.75, 0.75, 1, 0},
		0.25,
		Material{Plastic, Color{1, 1, 1}, 0.1, 1.45, false},
	})

	// BOTTOM
	list = append(list, Sphere{
		Tuple{0, -10000.5, -1, 0}, 10000,
		Material{Lambertian, Color{1, 1, 1}, 0, 0, true},
	})

	// TOP LIGHT
	// list = append(list, Sphere{
	// 	Tuple{0, 2 + 0.9, 1.2, 0}, 1,
	// 	Material{Emission, Color{5, 5, 5}, 0, 0, false},
	// })

	// TOP
	list = append(list, Sphere{
		Tuple{0, 10002, -1, 0}, 10000,
		Material{Emission, Color{1, 1, 1}, 0, 0, false},
	})

	// LEFT
	list = append(list, Sphere{
		Tuple{10002, 0, -1, 0}, 10000,
		Material{Lambertian, Color{1, 0.78, 0.23}, 0, 0, false},
	})

	// RIGHT
	list = append(list, Sphere{
		Tuple{-10002, 0, -1, 0}, 10000,
		Material{Lambertian, Color{0.32, 0.72, 1}, 0, 0, false},
	})

	// FRONT
	list = append(list, Sphere{
		Tuple{0, 0, -10001, 0}, 10000,
		Material{Lambertian, Color{1, 1, 1}, 0, 0, false},
	})

	// BACK
	list = append(list, Sphere{
		Tuple{0, 0, 10003, 0}, 10000,
		Material{Lambertian, Color{1, 1, 1}, 0, 0, false},
	})
	// list = append(list, Sphere{
	// 	Tuple{0.92602, -0.17328, -1, 0}, 0.239,
	// 	Material{Dielectric, Color{1, 0, 0}, 0, 1.333},
	// })
	// list = append(list, Sphere{
	// 	Tuple{-0.47711, -0.2623, -1.3436, 0}, 0.248,
	// 	Material{Lambertian, Color{1, 0.288, 0.302}, 0, 0},
	// })
	world := HittableList{list}

	// var wg sync.WaitGroup

	start := time.Now()

	for s := 0; s < samples; s++ {
		sample := time.Now()
		// wg.Add(1)
		// go func(s int) {
		for y := vsize - 1; y >= 0; y-- {
			for x := hsize - 1; x >= 0; x-- {
				col := Color{0, 0, 0}
				u := (float64(x) + RandFloat()) / float64(hsize)
				v := (float64(y) + RandFloat()) / float64(vsize)
				// r := Ray{origin, lower_left_corner.Add((horizontal.MulScalar(u)).Add(vertical.MulScalar(v)))}
				r := camera.getRay(u, v)

				col = color(r, &world, 0)

				// col = col.DivScalar(float64(samples))

				// p := r.Position(2.0)

				canvas[y*hsize+x] = canvas[y*hsize+x].Add(col)
			}
		}

		fmt.Printf("\r%.2f%% (% 3d/% 3d)", float64(s+1)/float64(samples)*100, s+1, samples)
		sampleTime := time.Since(sample)
		fmt.Printf(" % 15s/sample, % 15s sample time, ETA: % 15s", sampleTime, sampleTime/(vsize*hsize), sampleTime*(samples-time.Duration(s)-1))
		// wg.Done()
		// }(s)
	}

	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

	// wg.Wait()

	for y := 0; y < vsize; y++ {
		for x := 0; x < hsize; x++ {
			canvas[y*hsize+x] = canvas[y*hsize+x].DivScalar(float64(samples))
		}
	}

	fmt.Printf("\nSaving...\n")
	// filename := fmt.Sprintf("frame_%d.ppm", 0)
	filename := fmt.Sprintf("frame_%d.ppm", time.Now().UnixNano()/1e6)

	SaveImage(canvas, hsize, vsize, 255, filename)
}
