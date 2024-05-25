package main

import (
	"encoding/xml"
	"flag"
)

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []Cake   `xml:"cake" json:"cake"`
}

type Cake struct {
	Name        string        `xml:"name" json:"name"`
	Time        int           `xml:"stovetime" json:"time"`
	Ingredients []Ingregients `xml:"ingredients>item" json:"ingredients"`
}

type Ingregients struct {
	ItemName  string `xml:"itemname" json:"ingredient_name"`
	ItemCount int    `xml:"itemcount" json:"ingredient_count"`
	ItemUnit  string `xml:"itemunit,omitempty" json:"item_unit,omitempty"`
}

type DBReader interface {
	Read() error
}

var file string

func init() {
	flag.StringVar(&file, "old", "", "Old database file")
	flag.StringVar(&file, "new", "", "New database file")
}

func main() {
	flag.Parse()
}
