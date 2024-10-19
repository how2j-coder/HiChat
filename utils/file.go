package utils

import (
	"path"
	"runtime"
)

// GetRootPath 获取程序根目录
func GetRootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}
