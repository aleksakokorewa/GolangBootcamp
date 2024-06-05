package main

import (
	"GolangBootcamp/DatabaseReader/FSComparer/config"
	"GolangBootcamp/DatabaseReader/pkg/fscomparer"
	"fmt"
)

func main() {
	oldPath, newPath, err := config.CheckSnapshot()
	if err != nil {
		fmt.Println(err)
		return
	}
	er := fscomparer.FSCompare(oldPath, newPath)
	if er != nil {
		return
	}
}
