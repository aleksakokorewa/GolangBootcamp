package db_reader

import (
	"encoding/json"
	"fmt"
)

type JSONRecipe Recipe

func (j JSONRecipe) Read() error {
	data, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
