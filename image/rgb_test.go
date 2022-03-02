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

func TestRGB_GetBiggestColor(t *testing.T) {

	rgb := image.RGB{
		Red:   2,
		Green: 3,
		Blue:  4,
	}

	color, value := rgb.GetBiggestColor()
	assert.Equal(t, image.ColorBlue, color)
	assert.Equal(t, 4, value)
}

func TestRGB_SetColorValue(t *testing.T) {

	rgb := image.RGB{
		Red:   2,
		Green: 3,
		Blue:  4,
	}

	rgb.SetColorValue(image.ColorRed, 4)
	assert.Equal(t, 4, rgb.Red)
}

func TestRGBBin_GetBinaryByColor(t *testing.T) {
	rgb := image.RGB{
		Red:   2,
		Green: 3,
		Blue:  4,
	}

	rgbBin := rgb.DecimalToBinary()

	bin := rgbBin.GetBinaryByColor(image.ColorRed)
	assert.Equal(t, "00000010",bin)

}

func TestRGB_ShiftLeftColorValue(t *testing.T) {
	rgb := image.RGB{
		Red:   2,
		Green: 3,
		Blue:  4,
	}
	rgb.ShiftLeftColorValue()
	assert.Equal(t, 4,rgb.Red )
	assert.Equal(t, 6,rgb.Green )
	assert.Equal(t, 8,rgb.Blue )
}


func TestRGB_ShifRightColorValue(t *testing.T) {
	rgb := image.RGB{
		Red:   2,
		Green: 3,
		Blue:  4,
	}
	rgb.ShiftRightColorValue()
	assert.Equal(t, 1,rgb.Red )
	assert.Equal(t, 129,rgb.Green )
	assert.Equal(t, 2,rgb.Blue )
}