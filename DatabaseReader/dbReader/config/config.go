package config

import (
	"errors"
	"flag"
)

func CheckFile() (fileName string, err error) {
	var NoFilenameError = errors.New("передайте название файла через флаг -f")
	flag.StringVar(&fileName, "f", "", "input file")
	flag.Parse()
	if fileName == "" {
		return "", NoFilenameError
	}
	return fileName, err
}
