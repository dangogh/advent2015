package main

import (
	"./word"
	"fmt"
	"log"
	"os"
	"strconv"
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
	var num int
	if len(os.Args) > 1 {
		var err error
		if num, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatalf("Expected number of passwords to generate,  got %s", os.Args[1])
		}
	}
	pwch := PasswordGenerator("hepxcrrq")

	for pw := range pwch {
		acc := pw.Accept()
		if acc {
			fmt.Printf("%s is %v\n", pw, acc)
			num--
			if num <= 0 {
				break
			}
		}
	}
}
