package dbcomparer

import (
	"GolangBootcamp/DatabaseReader/pkg/dbreader"
	"fmt"
	"slices"
)

func GetCake(name string, recipe dbreader.Recipe) dbreader.Cake {
	for _, cake := range recipe.Cakes {
		if cake.Name == name {
			return cake
		}
	}
	return dbreader.Cake{}
}

func GetCakeName(cakes dbreader.Recipe) []string {
	var CakeNames []string
	for _, cake := range cakes.Cakes {
		CakeNames = append(CakeNames, cake.Name)
	}
	return CakeNames
}

func GetIngredients(name string, recipe dbreader.Cake) (ingredient dbreader.Ingredient) {
	for _, ingredient = range recipe.Ingredients {
		if name == ingredient.IngredientName {
			return ingredient
		}
	}
	return dbreader.Ingredient{}
}

func GetIngredient(cake dbreader.Cake) (ingredients []dbreader.Ingredient, name []string) {
	for _, ingredient := range cake.Ingredients {
		name = append(name, ingredient.IngredientName)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, name
}

func CompareUnitCount(oldCount dbreader.Ingredient, newCount dbreader.Ingredient, newCake string) {
	if oldCount.IngredientCount != newCount.IngredientCount {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newCount.IngredientName, newCake, newCount.IngredientCount, oldCount.IngredientCount)
	}
}

func CompareIngredientUnits(cakeOld dbreader.Ingredient, cakeNew dbreader.Ingredient, cakeNewName string) {
	switch {
	case cakeNew.IngredientUnit != cakeOld.IngredientUnit && cakeNew.IngredientUnit != "" && cakeOld.IngredientUnit != "":
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", cakeNew.IngredientName, cakeNewName, cakeNew.IngredientUnit, cakeOld.IngredientUnit)
	case cakeNew.IngredientUnit != "" && cakeOld.IngredientUnit == "":
		fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", cakeOld.IngredientUnit, cakeOld.IngredientName, cakeNewName)
	case cakeNew.IngredientUnit == "" && cakeOld.IngredientUnit != "":
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", cakeOld.IngredientUnit, cakeOld.IngredientName, cakeNewName)
	}
}

func CompareCakeName(oldCake, newCake []string) {
	for _, newCakes := range newCake {
		if !slices.Contains(oldCake, newCakes) {
			fmt.Printf("ADDED cake \"%s\"\n", newCakes)
		}
	}
	for _, oldCakes := range oldCake {
		if !slices.Contains(newCake, oldCakes) {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCakes)
		}
	}
}

func CompareTime(oldTime dbreader.Cake, newTime dbreader.Cake) {
	if oldTime.Time != newTime.Time {
		fmt.Printf("CHANGED cooking time - for cake \"%s\" - \"%s\" instead of \"%s\"\n", newTime.Name, newTime.Time, oldTime.Time)
	}
}

func Comparer(oldpath string, newpath string) {
	x := dbreader.XmlReader{}
	original, _ := x.ReadDB(oldpath)
	j := dbreader.JSONReader{}
	stolen, _ := j.ReadDB(newpath)
	oldNames := GetCakeName(original)
	stolenNames := GetCakeName(stolen)
	CompareCakeName(oldNames, stolenNames) //сравнение имен
	for _, cake := range oldNames {
		if slices.Contains(stolenNames, cake) {
			oldCake := GetCake(cake, original)
			newCake := GetCake(cake, stolen)
			CompareTime(oldCake, newCake)               //сравнение времени
			CompareIngredientsChanges(oldCake, newCake) //сравнение ингредиентов
		}
	}
}

func CompareIngredientsChanges(cakeOld dbreader.Cake, cakeNew dbreader.Cake) {
	oldIngred, oldNames := GetIngredient(cakeOld)
	newIngred, newNames := GetIngredient(cakeNew)
	for _, newIngredItem := range newIngred {
		if !slices.Contains(oldNames, newIngredItem.IngredientName) {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredItem.IngredientName, cakeNew.Name)
		} else {
			oldIng := GetIngredients(newIngredItem.IngredientName, cakeOld)
			CompareIngredientUnits(oldIng, newIngredItem, cakeNew.Name)
			CompareUnitCount(oldIng, newIngredItem, cakeNew.Name)
		}
	}
	for _, oldIngredItem := range oldIngred {
		if !slices.Contains(newNames, oldIngredItem.IngredientName) {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngredItem.IngredientName, cakeNew.Name)
		}
	}
}
