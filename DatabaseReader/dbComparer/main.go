package main

import (
	"GolangBootcamp/DatabaseReader/dbComparer/config"
	"GolangBootcamp/DatabaseReader/pkg/db_comparer"
	"GolangBootcamp/DatabaseReader/pkg/db_reader"
	"fmt"
)

func main() {
	oldDb, newDb, err := config.GetFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	oldcakes, _ := db_reader.CheckExtension(oldDb)
	newcakes, _ := db_reader.CheckExtension(newDb)
	db_comparer.Comparer(oldcakes, newcakes)
}
