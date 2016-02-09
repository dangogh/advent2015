package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Ingredient struct {
	Name                                            string
	Capacity, Durability, Flavor, Texture, Calories int
}

func ReadIngredients(r io.Reader) []Ingredient {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var ingr []Ingredient
	for s.Scan() {
		t := s.Text()
		var i Ingredient
		_, err := fmt.Sscanf(t, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&i.Name, &i.Capacity, &i.Durability, &i.Flavor, &i.Texture, &i.Calories)
		i.Name = strings.TrimRight(i.Name, `:`)
		if err != nil {
			fmt.Println("Line is ", t)
			panic(fmt.Sprintf("Error: %s:\n %s\n", err, t))
		}
		ingr = append(ingr, i)
	}
	return ingr
}

func main() {
	f, err := os.Open(os.Args[1])
	ingr := ReadIngredients(bufio.NewReader(f))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ingr: %+v\n", ingr)
}
