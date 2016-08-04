package main

import (
	"image"
	"image/color"
	"math"
)

// Builds a circle up to a given arctan
type arcSector struct {
	p     image.Point
	r     int
	angle float64
}

func (c *arcSector) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *arcSector) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *arcSector) At(x, y int) color.Color {
	dx := float64(x - c.p.X)
	dy := float64(y - c.p.Y)
	if dx*dx+dy*dy < float64(c.r*c.r) {
		angle := math.Atan2(dy, dx) + math.Pi/2
		if angle < 0 {
			angle += 2 * math.Pi
		}
		if angle < (c.angle * math.Pi / 180) {
			d := float64(c.r*c.r) - (dx*dx + dy*dy)
			alpha := uint8(255)

			// Basic edge smoothing
			for i := 255.0; i >= 0; i -= 5 {
				if d > 0 && d < i {
					alpha = uint8(i)
				}
			}

			return color.Alpha{alpha}
		}
	}
	return color.Alpha{0}
}
