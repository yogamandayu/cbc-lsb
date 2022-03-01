package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

// RootDir returns relative path of current project.
func RootDir() string {
	_, b, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Join(filepath.Dir(b), "..")
	}

	return ""
}

// ExtractMimeTypes returns slice of mime type.
func ExtractMimeTypes(mimeTypes string) []interface{} {
	sliceStr := strings.Split(mimeTypes, ",")
	sliceInterface := make([]interface{}, len(sliceStr))
	for i, v := range sliceStr {
		sliceInterface[i] = v
	}

	return sliceInterface
}
