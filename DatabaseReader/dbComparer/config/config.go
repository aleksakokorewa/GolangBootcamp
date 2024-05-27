package config

import (
	"errors"
	"flag"
)

func GetFile() (oldDB string, newDB string, err error) {
	var NoFilenameError = errors.New("передайте файлы в формате: --old filename --new filename")
	flag.StringVar(&oldDB, "old", "", "Old config file")
	flag.StringVar(&newDB, "new", "", "New config file")
	flag.Parse()
	if oldDB == "" || newDB == "" {
		return "", "", NoFilenameError
	}
	return oldDB, newDB, nil
}
