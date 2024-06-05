package config

import (
	"errors"
	"flag"
)

func CheckFile() (fileName string, err error) {
	var noFileError = errors.New("передайте название файла через флаг -f")
	flag.StringVar(&fileName, "f", "", "input file")
	flag.Parse()
	if fileName == "" {
		err = noFileError
	}
	return fileName, err
}
