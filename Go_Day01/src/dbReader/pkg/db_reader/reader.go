package db_reader

import "encoding/xml"

type DBReader interface {
	Read() error
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
