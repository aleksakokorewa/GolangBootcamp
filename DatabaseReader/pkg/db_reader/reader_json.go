package db_reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type JSONRecipe struct {
	cakes Recipe
}

type JSONconverter struct{}

func (j JSONRecipe) ReadDB(filename string) (Recipe, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return j.cakes, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &j.cakes)
	return j.cakes, nil
}

func (c JSONconverter) Converter(cake Recipe) {
	xmlCake, _ := xml.MarshalIndent(cake, " ", "    ")
	fmt.Println(string(xmlCake))
}
