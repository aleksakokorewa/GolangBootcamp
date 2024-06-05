package main

import (
	"GolangBootcamp/DatabaseReader/DBReader/config"
	"GolangBootcamp/DatabaseReader/pkg/dbreader"
	"fmt"
)

func main() {
	name, err := config.CheckFile()
	if err != nil {
		fmt.Println(err)
	}
	dbreader.CheckExtension(name)
}
