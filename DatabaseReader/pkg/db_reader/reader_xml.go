package db_reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type XMLRecipe struct {
	cake Recipe
}

type converterXML struct{}

func (x XMLRecipe) ReadDB(filename string) (Recipe, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return x.cake, err
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &x.cake)
	return x.cake, nil
}

func (x converterXML) Converter(cake Recipe) {
	jsonCake, _ := json.MarshalIndent(cake, " ", "    ")
	fmt.Println(string(jsonCake))
}
