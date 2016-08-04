package main

import (
	"bufio"
	"flag"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	fontFile := "/src/github.com/golang/freetype/testdata/luxisr.ttf"
	fontFile = os.Getenv("GOPATH") + fontFile

	fontFace, err := LoadFont(fontFile)
	if err != nil {
		log.Fatal("Could not load font: %s", err)
	}

	percentage := flag.Int("percentage", 360, "percentage to generate")
	flag.Parse()

	donut := NewDonut(
		float64(*percentage),
		300,
		color.RGBA{255, 255, 255, 255},
		color.RGBA{153, 194, 255, 255},
		color.RGBA{0, 102, 255, 255},
		color.RGBA{0, 0, 0, 255},
		fontFace, 12,
	)

	// Generate and draw the donut
	img := donut.Draw()
	file, err := os.Create("donut.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b := bufio.NewWriter(file)
	err = png.Encode(b, img)
	if err != nil {
		log.Fatal(err)
	}

	err = b.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
