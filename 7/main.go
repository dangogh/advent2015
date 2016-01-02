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

func (t token) String() string {
	s := "?"
	switch t.tokenType {
	case andTok:
		s = "*"
	case orTok:
		s = "+"
	case notTok:
		s = "!"
	case rshiftTok:
		s = ">>"
	case lshiftTok:
		s = "<<"
	case numTok:
		s = strconv.Itoa(t.value.(int))
	case identTok:
		s = t.value.(string)
	default:
		log.Fatalf("Unknown operator %d", t.tokenType)
	}
	return s
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

func (e expression) String() string {
	var s string
	operands := make([]string, len(e.operands))
	for _, a := range e.operands {
		operands = append(operands, a.String())
	}
	if len(operands) != 1 {
		s = strings.Join(operands, e.op.String())
	} else {
		s = e.op.String() + operands[0]
	}
	return s
}

var wireAssignments map[string]expression

func init() {
	wireAssignments = make(map[string]expression)
}

func parseWire(line string) {
	fmt.Println("Line: ", line)
	tokens := identifyTokens(strings.Split(line, " "))
	for i, t := range tokens {
		fmt.Printf("  %d - %#v\n", i, t)
	}

	var op token
	operands := []token{}
	for _, t := range tokens {
		switch t.tokenType {
		case numTok:
			fallthrough
		case identTok:
			operands = append(operands, t)
		case assignTok: // not needed
		default:
			op = t
		}
	}

	a, operands := operands[len(operands)-1], operands[:len(operands)-1]
	name := a.value.(string)
	wireAssignments[name] = expression{op, operands}
	//fmt.Printf("%s <- %s %v\n", name, op.String(), operands)
}

func main() {
	var reader *bufio.Reader
	infile, wires := os.Args[1], os.Args[2:]
	//if len(os.Args) == 1 {
	//	reader = bufio.NewReader(os.Stdin)
	//} else {
	f, err := os.Open(infile)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reading from %s\n", infile)
	reader = bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	//display := make(Display)
	for scanner.Scan() {
		line := scanner.Text()
		parseWire(line)
	}

	for _, w := range wires {
		if a, ok := wireAssignments[w]; !ok {
			log.Printf("No wire named %s\n", w)
		} else {
			fmt.Printf("%s <- %v\n", w, a)
		}
	}
}
