package configs

import (
	"path"
	"path/filepath"
	"runtime"
)

var basePath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath = path.Dir(currentFile)
}

func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(basePath, rel)
}
