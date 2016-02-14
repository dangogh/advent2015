package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ingredient [5]int

func readIngredients(r io.Reader) []ingredient {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var ingr []ingredient
	for s.Scan() {
		t := s.Text()
		var i ingredient
		var name string
		_, err := fmt.Sscanf(t, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name, &i[0], &i[1], &i[2], &i[3], &i[4])
		if err != nil {
			fmt.Println("Line is ", t)
			panic(fmt.Sprintf("Error: %s:\n %s\n", err, t))
		}
		ingr = append(ingr, i)
	}
	return ingr
}

func solve(ingr []ingredient) (int, []int) {
	var maxScore int
	var ans []int
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100-a; b++ {
			for c := 0; c <= 100-(a+b); c++ {
				d := 100 - (a + b + c)
				cl := ingr[0][4]*a + ingr[1][4]*b + ingr[2][4]*c + ingr[3][4]*d
				if cl != 500 {
					continue
				}

				ca := ingr[0][0]*a + ingr[1][0]*b + ingr[2][0]*c + ingr[3][0]*d
				du := ingr[0][1]*a + ingr[1][1]*b + ingr[2][1]*c + ingr[3][1]*d
				fl := ingr[0][2]*a + ingr[1][2]*b + ingr[2][2]*c + ingr[3][2]*d
				tx := ingr[0][3]*a + ingr[1][3]*b + ingr[2][3]*c + ingr[3][3]*d
				if ca <= 0 || du <= 0 || fl <= 0 || tx <= 0 {
					continue
				}
				v := ca * du * fl * tx
				if v > maxScore {
					maxScore = v
					ans = []int{a, b, c, d}
				}
			}
		}
	}
	return maxScore, ans
}

func main() {
	f, err := os.Open(os.Args[1])
	ingr := readIngredients(bufio.NewReader(f))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ingr: %+v\n", ingr)
	max, ans := solve(ingr)
	fmt.Printf("solved: %d;  %+v\n", max, ans)
}
