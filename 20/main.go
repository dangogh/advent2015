package main

import "fmt"

var houses []int

func main() {

	n := 29000000
	house := make([]int, n/10)
	for i := 1; i < n/10; i++ {
		for j := i; j < n/10; j += i {
			house[j] += i * 10
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
