package util_test

import (
	"cbc-lsb/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUtil_XOR(t *testing.T) {
	binary1 := "0"
	binary2 := "1"

	res1, err := util.XOR(binary1, binary1)
	require.NoError(t, err)
	assert.Equal(t, "0", res1)

	res2, err := util.XOR(binary2, binary2)
	require.NoError(t, err)
	assert.Equal(t, "0", res2)

	res3, err := util.XOR(binary1, binary2)
	require.NoError(t, err)
	assert.Equal(t, "1", res3)

	res4, err := util.XOR(binary2, binary1)
	require.NoError(t, err)
	assert.Equal(t, "1", res4)
}
