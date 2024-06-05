package main

import (
	"GolangBootcamp/DatabaseReader/DBComparer/config"
	"GolangBootcamp/DatabaseReader/pkg/dbcomparer"
	"fmt"
)

func main() {
	oldF, newF, err := config.CompareFiles()
	if err != nil {
		fmt.Println("Error:", err)
	}
	dbcomparer.Comparer(oldF, newF)
}
