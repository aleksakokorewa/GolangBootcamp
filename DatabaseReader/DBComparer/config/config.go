package config

import (
	"errors"
	"flag"
)

func CompareFiles() (oldFile string, newFile string, err error) {
	NoFileError := errors.New("передайте файлы в формате: --old filename --new filename")
	flag.StringVar(&oldFile, "old", "", "old filename")
	flag.StringVar(&newFile, "new", "", "new filename")
	flag.Parse()
	if oldFile == "" || newFile == "" {
		err = NoFileError
		return "", "", err
	}
	return oldFile, newFile, err
}
