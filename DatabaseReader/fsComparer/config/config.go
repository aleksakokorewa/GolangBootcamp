package config

import (
	"errors"
	"flag"
)

func GetSnapshot() (oldSnapshot, newSnapshot string, err error) {
	var noSnapshotError = errors.New("передайте файлы в формате: --old snapshot1.txt --new snapshot2.txt")

	flag.StringVar(&oldSnapshot, "old", "", "old snapshot")
	flag.StringVar(&newSnapshot, "new", "", "new snapshot")

	flag.Parse()
	if oldSnapshot == "" || newSnapshot == "" {
		err = noSnapshotError
	}
	return oldSnapshot, newSnapshot, err
}
