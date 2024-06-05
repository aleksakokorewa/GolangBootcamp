package dbreader

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
)

type DBReader interface {
	ReadDB(path string) (Recipe, error)
}

type Convertation interface {
	Convert(cake Recipe)
}

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

func CheckExtension(path string) {
	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		j := JSONReader{}
		res, err := j.ReadDB(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		c := JsonConvert{}
		c.Convert(res)
	case ".xml":
		x := XmlReader{}
		res, err := x.ReadDB(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		c := XmlConverter{}
		c.Convert(res)
	}
}
