package cbclsb_test

import (
	cbclsb "cbc-lsb"
	"cbc-lsb/image"
	"cbc-lsb/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestCBCLSB_Encryption(t *testing.T) {
	//file := util.RootDir() + "/test/file/lenna.jpg"
	file := util.RootDir() + "/test/file/Lenna_(test_image).png"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Key image not found, place your key image in same location with this file.")
	} else {
		fmt.Println("Key image found")
	}

	encryption := cbclsb.NewCBCLSBEncryption()
	keyRGB := &cbclsb.RGBKey{
		RGB: image.RGB{
			Red:   32,
			Green: 17,
			Blue:  123,
		},
	}
	plaintext := "Hello World!"
	ciphertext, err := encryption.Encrypt("Hello World!", keyRGB, f)
	require.NoError(t, err)

	//file = util.RootDir() + "/test/file/lenna.jpg"
	file = util.RootDir() + "/test/file/Lenna_(test_image).png"

	f, err = os.Open(file)
	if err != nil {
		log.Fatal("Key image not found, place your key image in same location with this file.")
	} else {
		fmt.Println("Key image found")
	}

	keyRGB = &cbclsb.RGBKey{
		RGB: image.RGB{
			Red:   32,
			Green: 17,
			Blue:  123,
		},
	}
	resPlaintext, err := encryption.Decrypt(ciphertext, keyRGB, f)
	require.NoError(t, err)

	assert.Equal(t, plaintext, resPlaintext)
}
