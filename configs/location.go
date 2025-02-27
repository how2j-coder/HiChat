package configs

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

var basePath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath = path.Dir(currentFile)
	fmt.Println("auto init")
}

func Path(rel string) string  {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(basePath, rel)
}
