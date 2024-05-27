package main

import (
	"GolangBootcamp/DatabaseReader/fsComparer/config"
	"GolangBootcamp/DatabaseReader/pkg/fs_comparer"
)

func main() {
	oldSnapshop, newSnapshot, err := config.GetSnapshot()
	if err != nil {
		return
	}
	fs_comparer.CompareInfo(oldSnapshop, newSnapshot)
}
