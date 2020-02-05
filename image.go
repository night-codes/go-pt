package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	PPM = iota
	PNG
)

// checks if there's an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveImage(canvas []Color, width, height, maxValue int, fileName string, extension int) {
	for y := height - 1; y >= 0; y-- {
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

	f, err := os.Create(fileName)
	check(err)

	if extension == PPM {
		w := bufio.NewWriter(f)
		defer w.Flush()

		_, err = fmt.Fprintf(w, "P3\n%d %d\n%d\n", width, height, maxValue)
		check(err)

		for y := height - 1; y >= 0; y-- {
			for x := 0; x < width; x++ {
				_, err := fmt.Fprintf(w, "%d %d %d ", int(math.Sqrt(canvas[y*width+x].r)*255), int(math.Sqrt(canvas[y*width+x].g)*255), int(math.Sqrt(canvas[y*width+x].b)*255))
				check(err)
			}
			_, err := fmt.Fprint(w, "\n")
			check(err)
		}
	} else if extension == PNG {
		image := image.NewRGBA(image.Rect(0, 0, width, height))
		for y := height - 1; y >= 0; y-- {
			for x := 0; x < width; x++ {
				image.SetRGBA(x, width-y, color.RGBA{uint8(math.Sqrt(canvas[y*width+x].r) * 255.9), uint8(math.Sqrt(canvas[y*width+x].g) * 255.9), uint8(math.Sqrt(canvas[y*width+x].b) * 255.9), 255})
			}
		}
		png.Encode(f, image)
	}
}
