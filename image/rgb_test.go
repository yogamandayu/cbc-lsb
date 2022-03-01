package image_test

import (
	"cbc-lsb/image"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRGB_DecimalToBinary(t *testing.T) {

	rgb := image.RGB{
		Red:   2,
		Green: 3,
		Blue:  4,
	}

	rgbBin := rgb.DecimalToBinary()
	assert.Equal(t, "00000010", rgbBin.Red)
	assert.Equal(t, "00000011", rgbBin.Green)
	assert.Equal(t, "00000100", rgbBin.Blue)
}
