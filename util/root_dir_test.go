package util_test

import (
	"cbc-lsb/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilRootDir(t *testing.T) {
	rootDir := util.RootDir()
	assert.NotEmpty(t, rootDir)
}
