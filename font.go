package main

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// Given a filepath to a font file (expecting a .ttf) will return the font.Font
// interface, or an error
func LoadFont(file string) (font.Face, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	f, err := truetype.Parse(b)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    14,
		Hinting: font.HintingNone,
		DPI:     72,
	})
	return face, nil
}
