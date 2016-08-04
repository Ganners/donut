package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Donut struct {
	angle      float64
	size       int
	background color.Color
	midground  color.Color
	foreground color.Color
	font       color.Color
	fontFace   font.Face
	fontSize   int
}

func NewDonut(
	angle float64,
	size int,
	background, midground, foreground, font color.Color,
	fontFace font.Face,
	fontSize int,
) *Donut {
	return &Donut{
		angle:      angle,
		size:       size,
		background: background,
		midground:  midground,
		foreground: foreground,
		font:       font,
		fontFace:   fontFace,
		fontSize:   fontSize,
	}
}

func (d *Donut) Draw() *image.RGBA {

	rgba := image.NewRGBA(image.Rect(0, 0, d.size, d.size))
	draw.Draw(
		rgba,
		rgba.Bounds(),
		image.NewUniform(d.background),
		image.ZP,
		draw.Src,
	)

	// Define a font drawer onto this rgba image
	drawer := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(d.font),
		Face: d.fontFace,
	}

	// Draw the outer circle in the midground colour
	draw.DrawMask(
		rgba,
		rgba.Bounds(),
		image.NewUniform(d.midground),
		image.ZP,
		&arcSector{
			image.Point{d.size / 2, d.size / 2},
			(d.size - 5) / 2,
			360,
		},
		image.ZP,
		draw.Over,
	)

	// Draw the outer angle circle in the foreground colour
	draw.DrawMask(
		rgba,
		rgba.Bounds(),
		image.NewUniform(d.foreground),
		image.ZP,
		&arcSector{
			image.Point{d.size / 2, d.size / 2},
			(d.size - 5) / 2, // Artificial margin of 5
			d.angle,
		},
		image.ZP,
		draw.Over,
	)

	// Draw the inner circle in the background colour
	draw.DrawMask(
		rgba,
		rgba.Bounds(),
		image.NewUniform(d.background),
		image.ZP,
		&arcSector{
			image.Point{d.size / 2, d.size / 2},
			d.size / 4,
			360,
		},
		image.ZP,
		draw.Over,
	)

	// Draw the angle in the middle
	stringToDraw := fmt.Sprintf("%.0f%%", d.angle)
	width := drawer.MeasureString(stringToDraw)
	height := fixed.I(d.fontSize)
	drawer.Dot.X = fixed.I(150) - (width / 2)
	drawer.Dot.Y = fixed.I(150) + (height / 2)
	drawer.DrawString(stringToDraw)

	return rgba
}
