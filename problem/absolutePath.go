package problem

import (
	"fmt"
	"strings"
)

func AbsolutePath(path string) string {
	paths := strings.Split(path, "/")
	var absolutePaths []string

	for _, path := range paths {
		if path == ".." {
			absolutePaths = absolutePaths[:len(absolutePaths)-1]
			continue
		}
		if path != "." && len(path) > 0 {
			absolutePaths = append(absolutePaths, path)
		}
	}

	var absolutePath string

	for _, path := range absolutePaths {
		absolutePath += fmt.Sprintf("%s/", path)
	}

	return "/" + absolutePath
}
