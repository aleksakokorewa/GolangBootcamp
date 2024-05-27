package db_reader

import (
	"GolangBootcamp/DatabaseReader/dbReader/config"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type DBReader interface {
	ReadDB(filename string) (Recipe, error)
}

type DBConverter interface {
	Converter(cakes Recipe)
}

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []Cake   `xml:"cake" json:"cake"`
}

type Cake struct {
	Name        string        `xml:"name" json:"name"`
	Time        string        `xml:"stovetime" json:"time"`
	Ingredients []Ingregients `xml:"ingredients>item" json:"ingredients"`
}

type Ingregients struct {
	ItemName  string `xml:"itemname" json:"ingredient_name"`
	ItemCount string `xml:"itemcount" json:"ingredient_count"`
	ItemUnit  string `xml:"itemunit,omitempty" json:"item_unit,omitempty"`
}

func Call() {
	file, err := config.CheckFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	ConvertFile(file)
}

func ConvertFile(file string) {
	extension := filepath.Ext(file)
	var converter DBConverter
	switch extension {
	case ".xml":
		converter = converterXML{}
	case ".json":
		converter = JSONconverter{}
	default:
		fmt.Printf("Unsupported file extension: %s\n", extension)
		os.Exit(1)
	}
	cake, err := CheckExtension(file)
	if err == nil {
		converter.Converter(cake)
	}
}

func CheckExtension(filename string) (Recipe, error) {
	extension := filepath.Ext(filename)
	var reader DBReader
	switch extension {
	case ".xml":
		reader = XMLRecipe{}
	case ".json":
		reader = JSONRecipe{}
	default:
		fmt.Printf("Unsupported file extension: %s\n", extension)
	}
	cake, err := reader.ReadDB(filename)
	return cake, err
}
