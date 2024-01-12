package file

import (
	"DistroJudge/log"
	"io/fs"
	"os"
)

func Read(path string) (string, error) {
	content, err := os.ReadFile("file.txt")
	if err != nil {
		log.Errorf("Read file %s error. err: %v", path, err)
	}
	return string(content), err
}

func Write(content, path string) error {
	err := os.WriteFile(path, []byte(content), fs.ModeAppend)
	return err
}
