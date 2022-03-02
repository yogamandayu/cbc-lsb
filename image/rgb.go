package image

import (
	"cbc-lsb/util"
	"strconv"
)

type RGBInterface interface {
	DecimalToBinary() *RGBBin
	GetBiggestColor() (Color, int)
	SetColorValue(color Color, value int)
	ShiftLeftColorValue()
}

type RGBBinInterface interface {
	GetBinaryByColor(color Color) string
}

// RGB is an RGB struct.
type RGB struct {
	Red   int
	Green int
	Blue  int
}

// RGBBin is an RGB binary form struct.
type RGBBin struct {
	Red   string
	Green string
	Blue  string
}

var _ RGBInterface = &RGB{}
var _ RGBBinInterface = &RGBBin{}

// DecimalToBinary convert rgb from decimal to binary.
func (r *RGB) DecimalToBinary() *RGBBin {
	return &RGBBin{
		Red:   util.BinLengthNormalizer(strconv.FormatInt(int64(r.Red), 2), 8),
		Green: util.BinLengthNormalizer(strconv.FormatInt(int64(r.Green), 2), 8),
		Blue:  util.BinLengthNormalizer(strconv.FormatInt(int64(r.Blue), 2), 8),
	}
}

// GetBiggestColor is to get the biggest color value between red, green, or blue.
func (r *RGB) GetBiggestColor() (Color, int) {
	size := r.Red
	color := ColorRed

	if r.Green > size {
		size = r.Green
		color = ColorGreen
	}

	if r.Blue > size {
		size = r.Blue
		color = ColorBlue
	}

	return color, size
}

// SetColorValue is to set/change r/g/b value by color.
func (r *RGB) SetColorValue(color Color, value int) {
	switch color {
	case ColorRed:
		r.Red = value
	case ColorGreen:
		r.Green = value
	case ColorBlue:
		r.Blue = value
	}
}

// ShiftLeftColorValue is to shift color value to the left by 1.
// Example : 01000001 -> 10000010
func (r *RGB) ShiftLeftColorValue() {
	redBin := util.BinLengthNormalizer(util.IntToBin(r.Red), 8)
	r.Red = util.BinToInt(redBin[1:(len(redBin))] + redBin[0:1])

	greenBin := util.BinLengthNormalizer(util.IntToBin(r.Green), 8)
	r.Green = util.BinToInt(greenBin[1:(len(greenBin))] + greenBin[0:1])

	blueBin := util.BinLengthNormalizer(util.IntToBin(r.Blue), 8)
	r.Blue = util.BinToInt(blueBin[1:(len(blueBin))] + blueBin[0:1])
}

// ShiftRightColorValue is to shift color value to the right by 1.
// Example : 01000001 -> 10100000
func (r *RGB) ShiftRightColorValue() {
	redBin := util.BinLengthNormalizer(util.IntToBin(r.Red), 8)
	r.Red = util.BinToInt(redBin[(len(redBin)-1):(len(redBin))] + redBin[0:(len(redBin)-1)])

	greenBin := util.BinLengthNormalizer(util.IntToBin(r.Green), 8)
	r.Green = util.BinToInt(greenBin[(len(greenBin)-1):(len(greenBin))] + greenBin[0:(len(greenBin)-1)])

	blueBin := util.BinLengthNormalizer(util.IntToBin(r.Blue), 8)
	r.Blue = util.BinToInt(blueBin[(len(blueBin)-1):(len(blueBin))] + blueBin[0:(len(blueBin)-1)])
}

// GetBinaryByColor return binary by color. Argument must be "R", "G", or "B".
func (r RGBBin) GetBinaryByColor(color Color) string {
	switch color {
	case ColorRed:
		return r.Red
	case ColorGreen:
		return r.Green
	case ColorBlue:
		return r.Blue
	}
	return ""
}
