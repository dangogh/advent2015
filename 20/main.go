package main

import "fmt"

var houses []int

func main() {

	n := 29000000
	house := make([]int, 1, n/10)
	for i := 1; i < n/10; i++ {
		var count int
		for j := i; count < 50; j += i {
			for len(house) <= j {
				house = append(house, 0)
			}
			house[j] += i * 11
			count++
		}
		if house[i] > n {
			fmt.Println(i, " has ", house[i])
			break
		}
	}
}

/*

house: 10
elf: 10
n: 100


*/
