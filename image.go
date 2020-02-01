package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// checks if there's an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// SaveImage writes canvas of `width` and `height` to .PPM file
func SaveImage(canvas []Color, width, height, maxValue int, fileName string) {
	// err := os.Remove(fileName)
	// check(err)

	f, err := os.Create(fileName)
	check(err)
	w := bufio.NewWriter(f)
	defer w.Flush()

	_, err = fmt.Fprintf(w, "P3\n%d %d\n%d\n", width, height, maxValue)
	check(err)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if canvas[y*width+x].r > 1.0 {
				canvas[y*width+x].r = 1.0
			}

			if canvas[y*width+x].g > 1.0 {
				canvas[y*width+x].g = 1.0
			}

			if canvas[y*width+x].b > 1.0 {
				canvas[y*width+x].b = 1.0
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, err := fmt.Fprintf(w, "%d %d %d ", int(math.Sqrt(canvas[y*width+x].r)*255), int(math.Sqrt(canvas[y*width+x].g)*255), int(math.Sqrt(canvas[y*width+x].b)*255))
			check(err)
		}
		_, err := fmt.Fprint(w, "\n")
		check(err)
	}
}
