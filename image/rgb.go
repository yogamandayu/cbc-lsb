package image

import (
	"cbc-lsb/util"
	"strconv"
)

type RGBInterface interface {
	DecimalToBinary() *RGBBin
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

func (r *RGB) DecimalToBinary() *RGBBin {
	return &RGBBin{
		Red:   util.BinLengthNormalizer(strconv.FormatInt(int64(r.Red), 2), 8),
		Green: util.BinLengthNormalizer(strconv.FormatInt(int64(r.Green), 2), 8),
		Blue:  util.BinLengthNormalizer(strconv.FormatInt(int64(r.Blue), 2), 8),
	}
}
