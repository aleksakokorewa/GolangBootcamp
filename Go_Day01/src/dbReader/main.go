package dbReader

import (
	"GolangBootcamp/Go_Day01/src/dbReader/pkg/db_reader"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path"
)

var file string

func init() {
	flag.StringVar(&file, "f", "", "input file")
}

func main() {
	flag.Parse()
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	exteinsion := path.Ext(file)
	switch exteinsion {
	case ".json":
		var recipe db_reader.JSONRecipe
		err = json.Unmarshal(data, &recipe)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		err = recipe.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	case ".xml":
		var recipe db_reader.XMLRecipe
		err = xml.Unmarshal(data, &recipe)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		err = recipe.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}
}
