package util

import (
	"fmt"
	"strconv"
	"strings"
)

// BinLengthNormalizer is a binary length normalizer.
func BinLengthNormalizer(binary string, length int) string {
	var str []string
	dif := length - len(binary)
	if dif <= 0 {
		return binary
	}
	for i := 0; i < dif; i++ {
		str = append(str, "0")
	}
	str = append(str, binary)
	s := strings.Join(str, "")
	return s
}

// StringToBin is to convert from string to binary.
func StringToBin(s string) string {
	var res string
	for _, c := range s {
		res += fmt.Sprintf("%08b", byte(c))
	}
	return res
}

// BinToString is to convert binary to string.
func BinToString(s string) string {
	var subStr, str string
	i := 0
	for _, r := range s {
		i++
		subStr += string(r)
		if i%8 == 0 {
			str += string(rune(BinToInt(subStr)))
			subStr = ""
			i = 0
		}
	}
	return str
}

// ChangeMSB is to change MSB(most significant bit) with new string.
func ChangeMSB(binary, msb string) (string, error) {
	if len(msb) > len(binary) {
		return "", fmt.Errorf("util.error.change_msb.invalid_length")
	}

	var str string

	for i, s := range binary {
		if i >= len(msb) {
			str += string(s)
		}
		if i == 0 {
			str += msb
		}
	}
	return str, nil
}

// BinToInt is to convert binary to integer.
func BinToInt(binary string) int {
	x, _ := strconv.ParseInt(binary, 2, 64)
	return int(x)
}

// IntToBin is to convert integer to binary. Return is not in byte.
// Example : 65 will return 1000001
func IntToBin(value int) string {
	return strconv.FormatInt(int64(value), 2)
}

// SplitBinTo2BitArray is to split binary into array of 2bit.
// Example: 00100101 will be converted into array of [00,10,01,01].
func SplitBinTo2BitArray(binary string) []string {
	var split []string
	var temp strings.Builder
	for i := 1; i <= len(binary); i++ {
		temp.WriteString(string(binary[i-1]))
		if i%2 == 0 {
			split = append(split, temp.String())
			temp.Reset()
		}
	}
	return split
}
