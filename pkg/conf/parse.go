package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
)

// Parse 将配置文件解析为 struct，包括 yaml、toml、json 等，如果 fs 不为空，则开启监听配置文件更改
func Parse(configFile string, obj interface{}, reloads ...func()) error {
	confFileAbs, err := filepath.Abs(configFile)
	if err != nil {
		return err
	}
	filePathStr, filename := filepath.Split(confFileAbs)
	ext := strings.TrimLeft(path.Ext(filename), ".")
	filename = strings.ReplaceAll(filename, "."+ext, "") // excluding suffix names
	viper.AddConfigPath(filePathStr)                     // path
	viper.SetConfigName(filename)                        // file name
	viper.SetConfigType(ext)                             // get the configuration type from the file name
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(obj)
	if err != nil {
		return err
	}

	if len(reloads) > 0 {
		watchConfig(obj, reloads...)
	}
	return nil
}

// Listening for profile updates
func watchConfig(obj interface{}, reloads ...func()) {
	viper.WatchConfig()

	// Note: OnConfigChange is called twice on Windows
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(obj)
		if err != nil {
			fmt.Println("viper.Unmarshal error: ", err)
		} else {
			for _, reload := range reloads {
				reload()
			}
		}
	})
}
