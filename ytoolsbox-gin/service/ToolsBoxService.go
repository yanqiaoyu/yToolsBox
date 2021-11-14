package service

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/viper"
)

// 清空存放脚本的目录
func RemoveAllScripts() error {

	dir, err := ioutil.ReadDir("./" + viper.GetString("tool.scriptBasePath"))
	if err != nil {
		return err
	}

	for _, d := range dir {
		os.RemoveAll(path.Join([]string{viper.GetString("tool.scriptBasePath"), d.Name()}...))
	}

	return nil
}
