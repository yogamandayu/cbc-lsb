package image_test

import (
	"cbc-lsb/image"
	"cbc-lsb/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestImageKind(t *testing.T) {

	t.Run("Test with png image. Expected no error", func(t *testing.T) {

		file := util.RootDir() + "/test/file/Lenna_(test_image).png"

		f, err := os.Open(file)
		if err != nil {
			log.Fatal("Key image not found, place your key.jpg in same location with this file.")
		} else {
			fmt.Println("Key image found")
		}

		img, err := image.NewImage(f)
		require.NoError(t, err)
		require.NotNil(t, img)
	})

	t.Run("Test with jpg image. Expected no error", func(t *testing.T) {

		file := util.RootDir() + "/test/file/lenna.jpg"

		f, err := os.Open(file)
		if err != nil {
			log.Fatal("Key image not found, place your key.jpg in same location with this file.")
		} else {
			fmt.Println("Key image found")
		}

		img, err := image.NewImage(f)
		require.NoError(t, err)
		require.NotNil(t, img)
	})
}

func TestImage_GetInitialValue(t *testing.T) {

	t.Run("Test with jpg image. Expected no error", func(t *testing.T) {

		file := util.RootDir() + "/test/file/lenna.jpg"

		f, err := os.Open(file)
		if err != nil {
			log.Fatal("Key image not found, place your key.jpg in same location with this file.")
		} else {
			fmt.Println("Key image found")
		}

		img, err := image.NewImage(f)
		require.NoError(t, err)
		require.NotNil(t, img)

		iv, err := img.GetInitialValue()
		assert.NoError(t, err)
		assert.NotNil(t, iv)
	})
}

func TestImage_GetPixel(t *testing.T) {

	t.Run("Test with jpg image. Expected no error", func(t *testing.T) {

		file := util.RootDir() + "/test/file/lenna.jpg"

		f, err := os.Open(file)
		if err != nil {
			log.Fatal("Key image not found, place your key.jpg in same location with this file.")
		} else {
			fmt.Println("Key image found")
		}

		img, err := image.NewImage(f)
		require.NoError(t, err)
		require.NotNil(t, img)

		rgb := img.GetPixel(image.Coordinate{
			PosX: 100,
			PosY: 100,
		})
		assert.NotNil(t, rgb)
	})

}
