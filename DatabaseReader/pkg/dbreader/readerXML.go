package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type XmlReader struct {
	cake Recipe
}

type XmlConverter struct{}

func (x XmlReader) ReadDB(path string) (Recipe, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file", err)
		return x.cake, err
	}
	defer file.Close()
	fileinfo, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file", err)
		return x.cake, err
	}
	err = xml.Unmarshal(fileinfo, &x.cake)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
		return x.cake, err
	}
	return x.cake, nil
}

func (x XmlConverter) Convert(cake Recipe) {
	JsonFile, _ := json.MarshalIndent(cake, "", "    ")
	fmt.Println(string(JsonFile))
}
