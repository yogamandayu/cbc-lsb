package cbclsb

import (
	"cbc-lsb/image"
	"cbc-lsb/util"
	"os"
)

type EncryptionInterface interface {
	Encrypt(plaintext string, keyRGB *RGBKey, keyImage *os.File) (string, error)
	Decrypt(ciphertext string, keyRGB *RGBKey, keyImage *os.File) (string, error)
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

/*
Encrypt is to encrypt plaintext with RGB as first key and image as second key.
*/
func (c *CBCLSB) Encrypt(plaintext string, keyRGB *RGBKey, keyImage *os.File) (string, error) {
	var err error
	var binCipher string

	c.Image, err = image.NewImage(keyImage)
	if err != nil {
		return "", err
	}

	currentCoordinate, err := c.Image.GetInitialValue()
	if err != nil {
		return "", err
	}

	pxBin := c.Image.GetPixel(currentCoordinate).DecimalToBinary()

	for _, subPt := range plaintext {
		sBin := util.StringToBin(string(subPt))
		arrBin := util.SplitBinTo2BitArray(sBin)

		var prefixOK bool
		var subCt string
		for i, value := range arrBin {
			if i == 0 {
				if value == "01" {
					prefixOK = true
				}

				subCt += value
				continue
			}
			var s string

			color, size := keyRGB.GetBiggestColor()
			rgbColorBin := util.BinLengthNormalizer(util.IntToBin(size), 8)

			ivColorBin := util.BinLengthNormalizer(pxBin.GetBinaryByColor(color), 8)
			if i == 1 && !prefixOK {
				s, err = util.XOR(value[1:2], rgbColorBin[7:8])
				if err != nil {
					return "", err
				}

				s, err = util.XOR(s, ivColorBin[7:8])
				if err != nil {
					return "", err
				}

				s = value[0:1] + s

				subCt += s

			} else {
				s, err = util.XOR(value, rgbColorBin[6:8])
				if err != nil {
					return "", err
				}

				s, err = util.XOR(s, ivColorBin[6:8])
				if err != nil {
					return "", err
				}

				subCt += s

			}

			rgbColorBin, err = util.ChangeMSB(rgbColorBin, s)
			if err != nil {
				return "", err
			}

			keyRGB.SetColorValue(color, util.BinToInt(rgbColorBin))
			keyRGB.ShiftLeftColorValue()

			currentCoordinate = c.Image.GetNextPixel(currentCoordinate, &keyRGB.RGB, color)

			pxBin = c.Image.GetPixel(currentCoordinate).DecimalToBinary()
		}
		binCipher += subCt
	}

	return util.BinToString(binCipher), nil
}

func (c *CBCLSB) Decrypt(ciphertext string, keyRGB *RGBKey, keyImage *os.File) (string, error) {
	var err error
	var binPlain string

	c.Image, err = image.NewImage(keyImage)
	if err != nil {
		return "", err
	}

	currentCoordinate, err := c.Image.GetInitialValue()
	if err != nil {
		return "", err
	}

	pxBin := c.Image.GetPixel(currentCoordinate).DecimalToBinary()

	for _, subCt := range ciphertext {
		sBin := util.StringToBin(string(subCt))
		arrBin := util.SplitBinTo2BitArray(sBin)

		var prefixOK bool
		var subPt string
		for i, value := range arrBin {
			if i == 0 {
				if value == "01" {
					prefixOK = true
				}

				subPt += value
				continue
			}
			var s string

			color, size := keyRGB.GetBiggestColor()
			rgbColorBin := util.BinLengthNormalizer(util.IntToBin(size), 8)

			ivColorBin := util.BinLengthNormalizer(pxBin.GetBinaryByColor(color), 8)
			if i == 1 && !prefixOK {
				s, err = util.XOR(value[1:2], ivColorBin[7:8])
				if err != nil {
					return "", err
				}

				s, err = util.XOR(s, rgbColorBin[7:8])
				if err != nil {
					return "", err
				}

				subPt += value[0:1] + s

			} else {
				s, err = util.XOR(value, ivColorBin[6:8])
				if err != nil {
					return "", err
				}

				s, err = util.XOR(s, rgbColorBin[6:8])
				if err != nil {
					return "", err
				}

				subPt += s

			}

			rgbColorBin, err = util.ChangeMSB(rgbColorBin, value)
			if err != nil {
				return "", err
			}

			keyRGB.SetColorValue(color, util.BinToInt(rgbColorBin))
			keyRGB.ShiftLeftColorValue()

			currentCoordinate = c.Image.GetNextPixel(currentCoordinate, &keyRGB.RGB, color)

			pxBin = c.Image.GetPixel(currentCoordinate).DecimalToBinary()
		}
		binPlain += subPt
	}

	return util.BinToString(binPlain), nil
}
