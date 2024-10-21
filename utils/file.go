package utils

import (
	"os"
	"path"
	"runtime"
)

// GetRootPath 获取程序根目录
func GetRootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
