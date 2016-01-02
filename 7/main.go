package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	numTok = iota
	identTok
	assignTok
	andTok
	orTok
	notTok
	lshiftTok
	rshiftTok
	// States
	exprState
	assignState
)

var commands map[string]int

func init() {
	commands = map[string]int{
		"AND":    andTok,
		"OR":     orTok,
		"NOT":    notTok,
		"LSHIFT": lshiftTok,
		"RSHIFT": rshiftTok,
		"->":     assignTok,
	}
}

type token struct {
	tokenType int
	value     interface{}
}

func identifyTokens(strs []string) []token {
	r := make([]token, 0, len(strs))
	var value interface{}

	for _, t := range strs {
		var tokType int
		var ok bool
		if tokType, ok = commands[t]; ok {
			value = nil
		} else if intval, err := strconv.Atoi(t); err == nil {
			tokType = numTok
			value = intval
		} else {
			// TODO  check if valid identifier
			tokType = identTok
			value = t
		}
		r = append(r, token{tokType, value})
	}
	return r
}

type expression struct {
	op       token
	operands []token
}

var assignments map[string]expression

func doCommand(line string) {
	fmt.Println("Line: ", line)
	tokens := identifyTokens(strings.Split(line, " "))
	// pop off last two tokens -- keep last one
	var op int
	operands := []token{}
	for _, t := range tokens {
		switch t.tokenType {
		case assignTok: // not needed
		case numTok:
			fallthrough
		case identTok:
			operands = append(operands, t)
		default:
			op = t.tokenType
		}
	}
	assigned, operands := operands[len(operands)-1], operands[:len(operands)-1]

	fmt.Printf("%s <- %v(%v)\n", assigned.value.(string), op, operands)
}

func main() {
	var reader *bufio.Reader
	if len(os.Args) == 1 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Reading from %s\n", os.Args[1])
		reader = bufio.NewReader(f)
	}
	scanner := bufio.NewScanner(reader)
	//display := make(Display)
	for scanner.Scan() {
		line := scanner.Text()
		doCommand(line)
	}
}
