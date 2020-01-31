package main

import (
	"fmt"
	"log"
	"time"
)

const (
	hsize = 480
	vsize = 320
)

func main() {
	canvas := make([]Color, vsize*hsize)
	start := time.Now()

	for y := 0; y < vsize; y++ {
		for x := 0; x < hsize; x++ {
			canvas[y*hsize+x] = Color{float64(x) / hsize, float64(y) / vsize, 255}
		}
	}

	fmt.Printf("\nSaving...\n")
	filename := fmt.Sprintf("frame_%d.ppm", 0)
	// filename := fmt.Sprintf("frame_%d.ppm", time.Now().UnixNano()/1e6)

	SaveImage(canvas, hsize, vsize, 255, filename)

	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)
}
