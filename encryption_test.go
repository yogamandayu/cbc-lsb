package cbclsb_test

import (
	cbclsb "cbc-lsb"
	"cbc-lsb/image"
	"cbc-lsb/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestCBCLSB_Encrypt(t *testing.T) {
	//file := util.RootDir() + "/test/file/lenna.jpg"
	file := util.RootDir() + "/test/file/Lenna_(test_image).png"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Key image not found, place your key.jpg in same location with this file.")
	} else {
		fmt.Println("Key image found")
	}

	encryption := cbclsb.NewCBCLSBEncryption()
	keyRGB := &cbclsb.RGBKey{
		RGB: image.RGB{
			Red:   35,
			Green: 17,
			Blue:  123,
		},
	}
	ciphertext, err := encryption.Encrypt("Hello World!", keyRGB, f)
	assert.NoError(t, err)
	log.Println(ciphertext)

	//file = util.RootDir() + "/test/file/lenna.jpg"
	file = util.RootDir() + "/test/file/Lenna_(test_image).png"

	f, err = os.Open(file)
	if err != nil {
		log.Fatal("Key image not found, place your key.jpg in same location with this file.")
	} else {
		fmt.Println("Key image found")
	}

	keyRGB = &cbclsb.RGBKey{
		RGB: image.RGB{
			Red:   35,
			Green: 17,
			Blue:  123,
		},
	}
	plaintext, err := encryption.Decrypt(ciphertext, keyRGB, f)
	assert.NoError(t, err)
	log.Println(plaintext)
}

func TestCBCLSB_Decrypt(t *testing.T) {
	// TODO: Test me please!
}
