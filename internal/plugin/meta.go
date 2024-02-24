package plugin

import (
	"path/filepath"
	"strings"
)

func Name(plugin string) string {
	return strings.TrimPrefix(filepath.Base(plugin), "servoc-ext_")
}
