package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type JSONReader struct {
	cake Recipe
}

type JsonConvert struct{}

func (j JSONReader) ReadDB(path string) (Recipe, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file", err)
		return j.cake, err
	}
	defer file.Close()
	fileinfo, err := io.ReadAll(file)
	err = json.Unmarshal(fileinfo, &j.cake)
	if err != nil {
		fmt.Println("Error reading file", err)
		return j.cake, err
	}
	return j.cake, nil
}

func (j JsonConvert) Convert(cake Recipe) {
	xmlCake, _ := xml.MarshalIndent(cake, "", "    ")
	fmt.Println(string(xmlCake))
}
