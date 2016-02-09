package main

import (
	"strings"
	"testing"
)

func Test2Ingredient(t *testing.T) {

	test_input := `
Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
`

	type Ingredient struct {
		Name       string
		Capacity   int
		Durability int
		Flavor     int
		Texture    int
		Calories   int
	}

	r := strings.NewReader(test_input)
	ingr := ReadIngredients(r)
	var cp, d, f, tx, cl int
	for _, i := range ingr {
		cp += i.Capacity
		d += i.Durability
		f += i.Flavor
		tx += i.Texture
		cl += i.Calories
	}

}
