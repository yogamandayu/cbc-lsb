package util_test

import (
	"cbc-lsb/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUtil_BinLengthNormalizer(t *testing.T) {
	bin := "01010"

	newBin := util.BinLengthNormalizer(bin, 8)
	assert.Equal(t, "00001010", newBin)
}

func TestUtil_StringToBin(t *testing.T) {
	s := "A"
	bin := util.StringToBin(s)
	assert.Equal(t, "01000001", bin)
}

func TestUtil_BinToString(t *testing.T) {
	bin := "01000001"
	s := util.BinToString(bin)
	assert.Equal(t, "A", s)
}

func TestUtil_ChangeMSB(t *testing.T) {
	bin := "01000001"
	msb := "10"
	s, err := util.ChangeMSB(bin, msb)
	require.NoError(t, err)
	assert.Equal(t, "10000001", s)

}

func TestUtil_SplitBinTo2BitArray(t *testing.T) {
	s := "A"
	bin := util.StringToBin(s)
	arr := util.SplitBinTo2BitArray(bin)
	assert.Equal(t, 4, len(arr))
	assert.Equal(t, "01", arr[0])
	assert.Equal(t, "00", arr[1])
	assert.Equal(t, "00", arr[2])
	assert.Equal(t, "01", arr[3])
}

func TestUtil_BinToInt(t *testing.T) {
	s := "A"
	bin := util.StringToBin(s)
	binInt := util.BinToInt(bin)
	assert.Equal(t, 65, binInt)
}

func TestUtil_IntToBin(t *testing.T) {
	n := 65
	binInt := util.IntToBin(n)
	assert.Equal(t, "1000001", binInt)
}
