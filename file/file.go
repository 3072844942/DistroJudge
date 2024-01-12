package file

import (
	"DistroJudge/log"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Path(path string) string {
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "/", "\\")
	}
	if runtime.GOOS == "linux" {
		path = strings.ReplaceAll(path, "\\", "/")
	}
	return path
}

func Read(path string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("文件不存在")
	}

	content, err := os.ReadFile("file.txt")
	if err != nil {
		log.Errorf("Read file %s error. err: %v", path, err)
	}
	return string(content), err
}

func Write(content, path string) error {
	// 获取文件所在目录
	dir := filepath.Dir(path)

	// 确保目录存在
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Errorf("创建目录 %s 错误. 错误信息: %v", dir, err)
		return err
	}

	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 文件不存在，创建文件
		file, err := os.Create(path)
		if err != nil {
			log.Errorf("创建文件 %s 错误. 错误信息: %v", path, err)
			return err
		}
		defer file.Close()
	}

	err := os.WriteFile(path, []byte(content), fs.ModeAppend)
	return err
}
