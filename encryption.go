package cbclsb

import (
	"cbc-lsb/image"
	"os"
)

type EncryptionInterface interface {
	Encrypt(plaintext string, key RGBKey, image *os.File) (string, error)
	Decrypt(ciphertext string, key RGBKey, image *os.File) (string, error)
}

// RGBKey is a key for encrypt and decrypt.
type RGBKey struct {
	image.RGB
}

// CBCLSB is a main struct for encryption.
type CBCLSB struct {
	Plaintext  string
	RGBKey     *RGBKey
	Ciphertext string
	Image      *image.Image
}

// NewCBCLSBEncryption is a constructor.
func NewCBCLSBEncryption() *CBCLSB {
	return &CBCLSB{}
}

var _ EncryptionInterface = &CBCLSB{}

func (c CBCLSB) Encrypt(plaintext string, key RGBKey, file *os.File) (string, error) {
	var err error

	c.Image, err = image.NewImage(file)
	if err != nil {
		return "", err
	}

	iv, err := c.Image.GetInitialValue()
	if err != nil {
		return "", err
	}

	ivRGB := c.Image.GetPixel(iv)
	_ = ivRGB.DecimalToBinary()

	return "", nil
}

func (c CBCLSB) Decrypt(ciphertext string, key RGBKey, file *os.File) (string, error) {
	var err error

	c.Image, err = image.NewImage(file)
	if err != nil {
		return "", err
	}

	iv, err := c.Image.GetInitialValue()
	if err != nil {
		return "", err
	}

	ivRGB := c.Image.GetPixel(iv)
	_ = ivRGB.DecimalToBinary()

	return "", nil
}
