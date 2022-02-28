package image

import (
	"image"
	"log"
	"math"
	"os"
	"sync"
)

type Interface interface {
	GetInitialValue() error
}

// RGB is an RGB struct.
type RGB struct {
	Red   int
	Green int
	Blue  int
}

type Image struct {
	File         *os.File
	InitialValue *Coordinate
	Properties   *Properties
}

type Coordinate struct {
	PosX int
	PosY int
}

type Properties struct {
	Pixels [][]*RGB
	Width  int
	Height int
}

func NewImage(file *os.File) (*Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]*RGB
	for y := 0; y < height; y++ {
		var row []*RGB
		for x := 0; x < width; x++ {
			row = append(row, rGBAToRGB(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return &Image{
		File: file,
		Properties: &Properties{
			Pixels: pixels,
			Width:  width,
			Height: height,
		},
	}, nil
}

var _ Interface = &Image{}

// rGBAToRGB img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rGBAToRGB(r, g, b, a uint32) *RGB {
	rx := int(r / 257)
	gx := int(g / 257)
	bx := int(b / 257)
	_ = int(a / 257)
	return &RGB{rx, gx, bx}
}

func (i Image) GetInitialValue() error {

	/*
		===============NOTES================

		ROW = HEIGHT = Coordinate Y = I
		COLUMN = WIDTH = Coordinate X = J

		====================================

	*/

	var rw, cw []RGB //row weight,col weight
	var mr, mc RGB   //mean row weight, mean col weight
	//var posX, posY int //x and y position for IV

	log.Println("Generate initialization vector.")
	log.Println("Generate time depends on the image size, please wait.")

	px := i.Properties.Pixels

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for x := 0; x < i.Properties.Height; x++ { //Get row weigth
			var cr RGB //count row
			for y := 0; y < i.Properties.Width; y++ {
				cr.Red += px[x][y].Red
				cr.Green += px[x][y].Green
				cr.Blue += px[x][y].Blue
			}
			cr.Red = cr.Red / i.Properties.Width
			cr.Green = cr.Green / i.Properties.Width
			cr.Blue = cr.Blue / i.Properties.Width
			rw = append(rw, cr)
		}
		for _, v := range rw { //sum all row RGB
			mr.Red += v.Red
			mr.Green += v.Green
			mr.Blue += v.Blue
		}
		mr.Red = mr.Red / len(rw) //mean all row RGB
		mr.Green = mr.Green / len(rw)
		mr.Blue = mr.Blue / len(rw)

		var min float64
		for y, v := range rw { // search the minimum difference between row weight and every pixel RGB, and the result is a Y location for IV
			sum := math.Abs(float64(v.Red-mr.Red)) + math.Abs(float64(v.Green-mr.Green)) + math.Abs(float64(v.Blue-mr.Blue))
			if y == 0 {
				min = sum
				i.InitialValue.PosY = y
			}
			if sum < min {
				min = sum
				i.InitialValue.PosY = y
			}
		}
		wg.Done()
	}()

	go func() {
		for x := 0; x < i.Properties.Width; x++ { //Get column weigth
			var cC RGB //count column
			for y := 0; y < i.Properties.Height; y++ {
				cC.Red += px[y][x].Red
				cC.Green += px[y][x].Green
				cC.Blue += px[y][x].Blue
			}
			cC.Red = cC.Red / i.Properties.Width
			cC.Green = cC.Green / i.Properties.Width
			cC.Blue = cC.Blue / i.Properties.Width
			cw = append(cw, cC)
		}
		for _, v := range cw { //sum all column RGB
			mc.Red += v.Red
			mc.Green += v.Green
			mc.Blue += v.Blue
		}
		mc.Red = mc.Red / len(cw) //mean all column RGB
		mc.Green = mc.Green / len(cw)
		mc.Blue = mc.Blue / len(cw)

		var min float64
		for x, v := range cw { // search the minimum difference between column weight and every pixel RGB, and the result is a X location for IV
			sum := math.Abs(float64(v.Red-mc.Red)) + math.Abs(float64(v.Green-mc.Green)) + math.Abs(float64(v.Blue-mc.Blue))
			if x == 0 {
				min = sum
				i.InitialValue.PosX = x
			}
			if sum < min {
				min = sum
				i.InitialValue.PosX = x
			}
		}
		wg.Done()
	}()
	wg.Wait()
	return nil
}
