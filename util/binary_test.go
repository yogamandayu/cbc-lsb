package util_test

import (
	"cbc-lsb/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinLengthNormalizer(t *testing.T) {
	bin := "01010"

	newBin := util.BinLengthNormalizer(bin, 8)
	assert.Equal(t, "00001010", newBin)
}
