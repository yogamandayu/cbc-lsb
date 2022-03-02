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

	iv, err := c.Image.GetInitialValue()
	if err != nil {
		return "", err
	}

	ivBin := c.Image.GetPixel(iv).DecimalToBinary()

	for _, subPt := range plaintext {
		sBin := util.StringToBin(string(subPt))
		arrBin := util.SplitBinTo2BitArray(sBin)

		// To make output ciphertext still readable, the first 3 bit mustn't change.
		// Example : 01010101 from this binary, the first 3 bit can't be change to 000,
		// it will become 00010101 or 21, it's not readable character.

		var prefixOK bool
		for i, value := range arrBin {
			if i == 0 {
				if value == "01" {
					prefixOK = true
				}
				binCipher += value
				continue
			}
			var subCt string

			color, size := keyRGB.GetBiggestColor()
			rgbColorBin := util.BinLengthNormalizer(util.IntToBin(size), 8)

			ivColorBin := util.BinLengthNormalizer(ivBin.GetBinaryByColor(color), 8)
			if i == 1 && !prefixOK {
				subCt, err = util.XOR(value[1:2], rgbColorBin[7:8])
				if err != nil {
					return "", err
				}

				subCt, err = util.XOR(subCt, ivColorBin[7:8])
				if err != nil {
					return "", err
				}

				subCt = value[0:1] + subCt

				binCipher += subCt

			} else {
				subCt, err = util.XOR(value, rgbColorBin[6:8])
				if err != nil {
					return "", err
				}

				subCt, err = util.XOR(subCt, ivColorBin[6:8])
				if err != nil {
					return "", err
				}

				binCipher += subCt

			}
			rgbColorBin, err = util.ChangeMSB(rgbColorBin, subCt)
			if err != nil {
				return "", err
			}
			keyRGB.SetColorValue(color, util.BinToInt(rgbColorBin))
			keyRGB.ShiftLeftColorValue()
		}
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

	iv, err := c.Image.GetInitialValue()
	if err != nil {
		return "", err
	}

	ivBin := c.Image.GetPixel(iv).DecimalToBinary()

	for _, subCt := range ciphertext {
		sBin := util.StringToBin(string(subCt))
		arrBin := util.SplitBinTo2BitArray(sBin)

		// To make output ciphertext still readable, the first 3 bit mustn't change.
		// Example : 01010101 from this binary, the first 3 bit can't be change to 000,
		// it will become 00010101 or 21, it's not readable character.

		var prefixOK bool
		for i, value := range arrBin {
			if i == 0 {
				if value == "01" {
					prefixOK = true
				}
				binPlain += value
				continue
			}
			var subPt string

			color, size := keyRGB.GetBiggestColor()
			rgbColorBin := util.BinLengthNormalizer(util.IntToBin(size), 8)

			ivColorBin := util.BinLengthNormalizer(ivBin.GetBinaryByColor(color), 8)
			if i == 1 && !prefixOK {
				subPt, err = util.XOR(value[1:2], rgbColorBin[7:8])
				if err != nil {
					return "", err
				}

				subPt, err = util.XOR(subPt, ivColorBin[7:8])
				if err != nil {
					return "", err
				}

				binPlain += value[0:1] + subPt

			} else {
				subPt, err = util.XOR(value, rgbColorBin[6:8])
				if err != nil {
					return "", err
				}

				subPt, err = util.XOR(subPt, ivColorBin[6:8])
				if err != nil {
					return "", err
				}

				binPlain += subPt

			}
			rgbColorBin, err = util.ChangeMSB(rgbColorBin, value)
			if err != nil {
				return "", err
			}
			keyRGB.SetColorValue(color, util.BinToInt(rgbColorBin))
			keyRGB.ShiftLeftColorValue()
		}
	}

	return util.BinToString(binPlain), nil
}
