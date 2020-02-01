package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	hsize   = 800
	vsize   = 400
	samples = 16
)

func color(r Ray, world *HittableList, depth int) Color {
	rec := HitRecord{}
	if world.hit(r, Epsilon, math.MaxFloat64, &rec) {
		var attenuation Color
		var scattered Ray
		if depth < 50 && rec.material.Scatter(r, rec, &attenuation, &scattered) {
			return attenuation.Mul(color(scattered, world, depth+1))
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
	temp := make([][]Color, samples)
	for i := 0; i < samples; i++ {
		canvas := make([]Color, vsize*hsize)
		temp[i] = canvas
	}
	fmt.Printf("Rendering %dx%d at %d samples\n", hsize, vsize, samples)
	start := time.Now()

	lower_left_corner := Tuple{-2.0, -1.0, -1.0, 1}
	horizontal := Tuple{4.0, 0.0, 0.0, 1}
	vertical := Tuple{0.0, 2.0, 0.0, 1}
	origin := Tuple{0.0, 0.0, 0.0, 1}

	list := []Sphere{}
	list = append(list, Sphere{
		Tuple{0, 0, -2, 0}, 0.5,
		Material{Metal, Color{0.8, 0.186, 0.082}, 0.5, 0},
	})
	list = append(list, Sphere{
		Tuple{1, 0, -2, 0}, 0.5,
		Material{Metal, Color{0.9, 0.736, 0.356}, 0.5, 0},
	})
	list = append(list, Sphere{
		Tuple{-1, 0, -2, 0}, 0.5,
		Material{Metal, Color{0.95, 0.95, 0.95}, 0.05, 0},
	})
	list = append(list, Sphere{
		Tuple{0, -10000.5, -1, 0}, 10000,
		Material{Lambertian, Color{0.5, 0.5, 0.5}, 0, 0},
	})
	list = append(list, Sphere{
		Tuple{0.8, 0, -1, 0}, 0.5,
		Material{Dielectric, Color{1, 0, 0}, 0, 1.333},
	})
	world := HittableList{list}

	// var wg sync.WaitGroup

	for s := 0; s < samples; s++ {
		// wg.Add(1)
		// go func(s int) {
		canvas := temp[s]
		for y := vsize - 1; y >= 0; y-- {
			for x := 0; x < hsize; x++ {
				col := Color{0, 0, 0}
				u := (float64(x) + RandFloat()) / float64(hsize)
				v := (float64(y) + RandFloat()) / float64(vsize)
				r := Ray{origin, lower_left_corner.Add((horizontal.MulScalar(u)).Add(vertical.MulScalar(v)))}

				col = color(r, &world, 0)

				// col = col.DivScalar(float64(samples))

				// p := r.Position(2.0)

				canvas[y*hsize+x] = col
			}
		}
		fmt.Printf("\r%.2f%% (% 3d/% 3d)", float64(s+1)/float64(samples)*100, s+1, samples)
		// wg.Done()
		// }(s)
	}

	// wg.Wait()

	canvas := make([]Color, vsize*hsize)

	for s := 0; s < samples; s++ {
		for y := vsize - 1; y >= 0; y-- {
			for x := 0; x < hsize; x++ {
				canvas[y*hsize+x] = canvas[y*hsize+x].Add(temp[s][y*hsize+x])
			}
		}
	}

	for y := vsize - 1; y >= 0; y-- {
		for x := 0; x < hsize; x++ {
			canvas[y*hsize+x] = canvas[y*hsize+x].DivScalar(float64(samples))
		}
	}

	fmt.Printf("\nSaving...\n")
	filename := fmt.Sprintf("frame_%d.ppm", 0)
	// filename := fmt.Sprintf("frame_%d.ppm", time.Now().UnixNano()/1e6)

	SaveImage(canvas, hsize, vsize, 255, filename)

	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)
}
