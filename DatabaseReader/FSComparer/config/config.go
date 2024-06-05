package config

import (
	"errors"
	"flag"
)

func CheckSnapshot() (OldPath string, NewPath string, err error) {
	flag.StringVar(&OldPath, "old", "", "old path")
	flag.StringVar(&NewPath, "new", "", "new path")
	flag.Parse()
	var NoFileError = errors.New("передайте файлы в формате: --old snapshot1.txt --new snapshot2.txt")
	if OldPath == "" || NewPath == "" {
		err := NoFileError
		return "", "", err
	}
	return OldPath, NewPath, nil
}
