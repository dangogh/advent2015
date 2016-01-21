package main

import (
	"./word"
	"fmt"
)

func PasswordGenerator(initialPw word.Word) chan word.Word {
	pwch := make(chan word.Word)

	go func() {
		w := initialPw
		for {
			w = w.Next()
			pwch <- w
		}

	}()
	return pwch
}

func main() {
	pwch := PasswordGenerator("abcdefzz")

	for pw := range pwch {
		acc := pw.Accept()
		if acc {
			fmt.Printf("%s is %v\n", pw, acc)
		}
		break
	}
}
