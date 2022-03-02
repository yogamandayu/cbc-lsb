package util

import "fmt"

// XOR is a digital gate of XOR.
func XOR(binary1 string, binary2 string) (string, error) {
	if len(binary1) != len(binary2) {
		return "", fmt.Errorf("util.error.xor.different_length")
	}
	return BinLengthNormalizer(IntToBin(BinToInt(binary1)^BinToInt(binary2)), len(binary1)), nil
}
