package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	hsize   = 128
	vsize   = 128
	samples = 64
	depth   = 8
)

func color(r Ray, world *HittableList, d int, generator rand.Rand) Color {
	rec := HitRecord{}
	if world.hit(r, Epsilon, math.MaxFloat64, &rec) {
		var attenuation Color
		var scattered Ray
		if d < depth && rec.material.Scatter(r, rec, &attenuation, &scattered, generator) {
			if rec.material.material == Emission {
				return rec.material.albedo
			} else {
				return attenuation.Mul(color(scattered, world, d+1, generator))
			}
		} else {
			return Color{0, 0, 0}
		}
	} else {
		// unit_direction := r.direction.Normalize()
		// t := 0.5 * (unit_direction.y + 1.0)
		// return Color{1.0, 1.0, 1.0}.MulScalar(1.0 - t).Add(Color{0.5, 0.7, 1.0}.MulScalar(t))
		return Color{1, 1, 1}
	}
}

func main() {
	listSpheres := []Sphere{}
	listTriangles := []Triangle{}

	cameraPosition := Tuple{-200, 2, 200, 0}
	cameraDirection := Tuple{20, 1, 1, 0}
	focusDistance := Tuple{0, 0, 0, 0}.Subtract(cameraPosition).Magnitude()
	// focusDistance := cameraDirection.Subtract(cameraPosition).Magnitude()
	camera := getCamera(cameraPosition, cameraDirection, Tuple{0, 1, 0, 0}, 75, float64(hsize)/float64(vsize), 0, focusDistance)

	vertices := []Tuple{}
	vertNormals := []Tuple{}
	faceVerts := [][3]Tuple{}
	faceNormals := [][3]Tuple{}

	file, err := os.Open("softball.obj")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		if text[0] == "v" {
			x, _ := strconv.ParseFloat(text[1], 64)
			y, _ := strconv.ParseFloat(text[2], 64)
			z, _ := strconv.ParseFloat(text[3], 64)
			vertices = append(vertices, Tuple{
				x, y, z, 0,
			})
		} else if text[0] == "vn" {
			x, _ := strconv.ParseFloat(text[1], 64)
			y, _ := strconv.ParseFloat(text[2], 64)
			z, _ := strconv.ParseFloat(text[3], 64)
			vertNormals = append(vertNormals, Tuple{
				x, y, z, 0,
			})
		} else if text[0] == "f" {
			values1 := strings.Split(text[1], "/")
			values2 := strings.Split(text[2], "/")
			values3 := strings.Split(text[3], "/")

			v1, _ := strconv.Atoi(values1[0])
			v2, _ := strconv.Atoi(values2[0])
			v3, _ := strconv.Atoi(values3[0])

			vn1, _ := strconv.Atoi(values1[2])
			vn2, _ := strconv.Atoi(values2[2])
			vn3, _ := strconv.Atoi(values3[2])

			faceVerts = append(faceVerts, [3]Tuple{
				vertices[v1-1], vertices[v2-1], vertices[v3-1],
			})

			faceNormals = append(faceNormals, [3]Tuple{
				vertNormals[vn1-1], vertNormals[vn2-1], vertNormals[vn3-1],
			})
		}
	}

	for i := 0; i < len(faceVerts); i++ {
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				faceVerts[i][0],
				faceVerts[i][1],
				faceVerts[i][2],
			},
			Material{Plastic, Color{0.1, 0.1, 0.1}, 0, 1.45, false},
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// BOTTOM
	listSpheres = append(listSpheres, Sphere{
		Tuple{0, -1000100, -1, 0}, 1000000,
		Material{Lambertian, Color{0.14, 0.74, 0.35}, 0, 1.45, true},
	})

	/*
		// listSpheres = append(listSpheres, Sphere{
		// 	Tuple{0, 0, 0, 0},
		// 	0.25,
		// 	Material{Dielectric, Color{1, 1, 1}, 0, 1.45, false},
		// })

		// BOTTOM SQUARE
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 0, 0, 0},
				Tuple{0, 0, 0, 0},
				Tuple{1, 0, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{0, 0, 0, 0},
				Tuple{0, 0, 1, 0},
				Tuple{1, 0, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		// TOP SQUARE
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 1, 0, 0},
				Tuple{0, 1, 0, 0},
				Tuple{1, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{0, 1, 0, 0},
				Tuple{0, 1, 1, 0},
				Tuple{1, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		// BACK SQUARE
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 0, 0, 0},
				Tuple{0, 0, 0, 0},
				Tuple{1, 1, 0, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 1, 0, 0},
				Tuple{0, 0, 0, 0},
				Tuple{0, 1, 0, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		// FRONT SQUARE
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 0, 1, 0},
				Tuple{0, 0, 1, 0},
				Tuple{1, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 1, 1, 0},
				Tuple{0, 0, 1, 0},
				Tuple{0, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		// RIGHT SQUARE
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{0, 0, 0, 0},
				Tuple{0, 0, 1, 0},
				Tuple{0, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{0, 1, 0, 0},
				Tuple{0, 0, 0, 0},
				Tuple{0, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		// LEFT SQUARE
		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 0, 1, 0},
				Tuple{1, 0, 0, 0},
				Tuple{1, 1, 1, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		listTriangles = append(listTriangles, Triangle{
			TrianglePosition{
				Tuple{1, 1, 0, 0},
				Tuple{1, 1, 1, 0},
				Tuple{1, 0, 0, 0},
			},
			Material{Metal, Color{1, 1, 1}, 0.4, 1.45, false},
		})

		// listTriangles = append(listTriangles, Triangle{
		// 	TrianglePosition{
		// 		Tuple{0, 0.25, 0, 0},
		// 		Tuple{1, 0.25, 0, 0},
		// 		Tuple{1, 0.25, 1, 0},
		// 	},
		// 	Material{Dielectric, Color{1, 0.5, 0.5}, 0, 1.45, false},
		// })

		// listTriangles = append(listTriangles, Triangle{
		// 	TrianglePosition{
		// 		Tuple{0, 1, -1, 0},
		// 		Tuple{0, -1, -1, 0},
		// 		Tuple{0, 0, 0, 0},
		// 	},
		// 	Material{Lambertian, Color{0.9, 0.1, 0.1}, 0.1, 1.45, false},
		// })

		// // TOP LIGHT
		// listSpheres = append(listSpheres, Sphere{
		// 	Tuple{0, 1.7, 0, 0}, 0.5,
		// 	Material{Emission, Color{5, 5, 5}, 0, 0, false},
		// })

		// TOP
		listSpheres = append(listSpheres, Sphere{
			Tuple{0, 10002, -1, 0}, 10000,
			Material{Emission, Color{1, 1, 1}, 0, 0, false},
		})

		// LEFT
		listSpheres = append(listSpheres, Sphere{
			Tuple{10002, 0, -1, 0}, 10000,
			Material{Lambertian, Color{1, 0.78, 0.23}, 0, 0, false},
		})

		// RIGHT
		listSpheres = append(listSpheres, Sphere{
			Tuple{-10002, 0, -1, 0}, 10000,
			Material{Lambertian, Color{0.32, 0.72, 1}, 0, 0, false},
		})

		// FRONT
		listSpheres = append(listSpheres, Sphere{
			Tuple{0, 0, -10001, 0}, 10000,
			Material{Lambertian, Color{1, 1, 1}, 0, 0, false},
		})

		// BACK
		listSpheres = append(listSpheres, Sphere{
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
		// })*/
	world := HittableList{listSpheres, listTriangles}

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	buf := make([][]Color, cpus)

	for i := 0; i < cpus; i++ {
		buf[i] = make([]Color, vsize*hsize)
	}

	ch := make(chan int, cpus)

	canvas := make([]Color, vsize*hsize)

	start := time.Now()

	samplesCPU := samples / cpus

	if samples < cpus {
		cpus = samples
		samplesCPU = samples
	}

	doneSamples := 0

	fmt.Printf("Rendering %dx%d at %d samples on %d cores\n", hsize, vsize, samples, cpus)

	for i := 0; i < cpus; i++ {
		go func(i int) {
			for s := 0; s < samplesCPU; s++ {
				source := rand.NewSource(time.Now().UnixNano())
				generator := rand.New(source)
				sample := time.Now()
				for y := vsize - 1; y >= 0; y-- {
					for x := 0; x < hsize; x++ {
						col := Color{0, 0, 0}
						u := (float64(x) + RandFloat(*generator)) / float64(hsize)
						v := (float64(y) + RandFloat(*generator)) / float64(vsize)
						r := camera.getRay(u, v, *generator)

						col = color(r, &world, 0, *generator)

						buf[i][y*hsize+x] = buf[i][y*hsize+x].Add(col)
					}
				}

				doneSamples++
				fmt.Printf("\r%.2f%% (% 3d/% 3d)", float64(doneSamples)/float64(samples)*100, doneSamples, samples)
				sampleTime := time.Since(sample)
				fmt.Printf(" % 15s/sample, % 15s sample time, ETA: % 15s", sampleTime, sampleTime/(vsize*hsize), sampleTime*(time.Duration(samples)-time.Duration(doneSamples))/time.Duration(cpus))
			}
			ch <- 1
		}(i)
	}

	for i := 0; i < cpus; i++ {
		<-ch
	}
	close(ch)

	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)
	for i := 0; i < cpus; i++ {
		for y := 0; y < vsize; y++ {
			for x := 0; x < hsize; x++ {
				canvas[y*hsize+x] = canvas[y*hsize+x].Add(buf[i][y*hsize+x])
			}
		}
	}

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
