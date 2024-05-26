package db_reader

import (
	"encoding/xml"
	"fmt"
)

type XMLRecipe Recipe

func (x XMLRecipe) Read() error {
	data, err := xml.MarshalIndent(x, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
