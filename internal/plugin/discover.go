package plugin

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Discover() []string {
	path := os.ExpandEnv(os.Getenv("PATH"))
	dirs := strings.Split(path, ":")
	extensions := make(map[string]struct{})
	for _, dir := range dirs {
		filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if d == nil {
				return nil
			}

			if path != dir && d.IsDir() {
				return filepath.SkipDir
			}
			if strings.HasPrefix(d.Name(), "servoc-ext_") {
				extensions[d.Name()] = struct{}{}
			}
			return nil
		})
	}
	ext := make([]string, len(extensions))
	i := 0
	for n := range extensions {
		ext[i] = n
		i++
	}
	return ext
}
