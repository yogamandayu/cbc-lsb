package util

import "strings"

func BinLengthNormalizer(binary string, length int) string {
	var str []string
	dif := length - len(binary)
	if dif == 0 {
		return binary
	}
	for i := 0; i < dif; i++ {
		str = append(str, "0")
	}
	str = append(str, binary)
	s := strings.Join(str, "")
	return s
}
