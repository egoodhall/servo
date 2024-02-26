package plugin

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Name(plugin string) string {
	return strings.TrimPrefix(filepath.Base(plugin), "servoc-ext_")
}

func OptionName(plugin, name string) string {
	return fmt.Sprintf("%s.%s", Name(plugin), name)
}
