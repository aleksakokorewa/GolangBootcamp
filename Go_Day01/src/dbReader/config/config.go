package config

import (
	"errors"
	"flag"
)

var NoFile = errors.New("передайте название файла через флаг -f")

func checkFile() (fileName string, err error) {
	flag.StringVar(&fileName, "f", "", "input file")
	flag.Parse()
	if fileName == "" {
		return "", NoFile
	}
	return fileName, err
}
