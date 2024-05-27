package db_comparer

import (
	"GolangBootcamp/DatabaseReader/pkg/db_reader"
	"fmt"
	"slices"
)

func Comparer(oldCakes db_reader.Recipe, newCakes db_reader.Recipe) {
	oldNames := getNames(oldCakes)
	compareNames(oldCakes, newCakes)
	for _, newCake := range newCakes.Cake {
		if slices.Contains(oldNames, newCake.Name) {
			oldCake := getCake(newCake.Name, oldCakes)
			compareTime(oldCake, newCake)
			compareIngredientsChanges(oldCake, newCake)
		}
	}
}

func getCake(name string, recipe db_reader.Recipe) (cake db_reader.Cake) {
	for _, recipeCake := range recipe.Cake {
		if recipeCake.Name == name {
			cake = recipeCake
		}
	}
	return
}

func getNames(recipe db_reader.Recipe) (names []string) {
	for _, cakes := range recipe.Cake {
		names = append(names, cakes.Name)
	}
	return names
}

func compareNames(oldDB db_reader.Recipe, newDB db_reader.Recipe) {
	oldNames := getNames(oldDB)
	newNames := getNames(newDB)
	for _, cake := range newDB.Cake {
		if !slices.Contains(oldNames, cake.Name) {
			fmt.Printf("ADDED cake \"%s\"\n", cake.Name)
		}
	}
	for _, cake := range oldDB.Cake {
		if !slices.Contains(newNames, cake.Name) {
			fmt.Printf("REMOVED cake \"%s\"\n", cake.Name)
		}
	}
}

func compareTime(oldDB db_reader.Cake, newDB db_reader.Cake) {
	if oldDB.Time != newDB.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", newDB.Name, newDB.Time, oldDB.Time)
	}
}

func getIngredients(cake db_reader.Cake) (ingredients []db_reader.Ingregients, names []string) {
	for _, ingredient := range cake.Ingredients {
		names = append(names, ingredient.ItemName)
		ingredients = append(ingredients, ingredient)
	}
	return
}

func getIngredient(name string, ingredients []db_reader.Ingregients) (ingredient db_reader.Ingregients) {
	for _, value := range ingredients {
		if value.ItemName == name {
			ingredient = value
			return
		}
	}
	return
}

func compareIngredientsCount(oldIngredient db_reader.Ingregients, newIngredient db_reader.Ingregients, newCakeName string) {
	switch {
	case oldIngredient.ItemCount != newIngredient.ItemCount:
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.ItemName, newCakeName, newIngredient.ItemCount, oldIngredient.ItemCount)
	case oldIngredient.ItemCount != "" && newIngredient.ItemCount == "":
		fmt.Printf("REMOVED unit count for ingredient \"%s\" for cake \"%s\"\n", newIngredient.ItemName, newCakeName)
	case oldIngredient.ItemCount == "" && newIngredient.ItemCount != "":
		fmt.Printf("ADDED unit count for ingredient \"%s\" for cake \"%s\"\n", newIngredient.ItemName, newCakeName)
	}
}

func compareIngredientsUnit(oldIngredient db_reader.Ingregients, newIngredient db_reader.Ingregients, newCakeName string) {
	switch {
	case oldIngredient.ItemUnit != newIngredient.ItemUnit:
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.ItemName, newCakeName, newIngredient.ItemUnit, oldIngredient.ItemUnit)
	case oldIngredient.ItemUnit != "" && newIngredient.ItemUnit == "":
		fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n", newIngredient.ItemName, newCakeName)
	case oldIngredient.ItemUnit == "" && newIngredient.ItemUnit != "":
		fmt.Printf("ADDED unit for ingredient \"%s\" for cake \"%s\"\n", newIngredient.ItemName, newCakeName)
	}
}

func compareIngredientsChanges(oldCake db_reader.Cake, newCake db_reader.Cake) {
	newIngred, newIngredNames := getIngredients(newCake)
	oldIngred, oldIngredNames := getIngredients(oldCake)

	for _, ingredients := range newIngred {
		if !slices.Contains(oldIngredNames, ingredients.ItemName) {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", ingredients.ItemName, newCake.Name)
		} else {
			oldIngredient := getIngredient(ingredients.ItemName, oldIngred)
			compareIngredientsUnit(oldIngredient, ingredients, newCake.Name)
			compareIngredientsCount(oldIngredient, ingredients, newCake.Name)
		}
	}
	for _, ingredients := range oldIngred {
		if !slices.Contains(newIngredNames, ingredients.ItemName) {
			fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", ingredients.ItemName, newCake.Name)
		}
	}
}
