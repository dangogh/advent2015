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
	assignTok = iota
	andTok
	orTok
	notTok
	lshiftTok
	rshiftTok
	numTok
	identTok
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
	case assignTok:
		s = "<-"
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
		//log.Printf("value of num: %d\n", t.value.(uint16))
		s = strconv.Itoa(int(t.value.(uint16)))
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
			value = uint16(intval)
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
	resolved bool
	result   uint16
}

func (e expression) String() string {
	var s string
	operands := make([]string, 0, len(e.operands))
	for _, a := range e.operands {
		operands = append(operands, a.String())
	}
	s = strings.Join(operands, e.op.String())
	return s
}

var wireAssignments map[string]*expression

func init() {
	wireAssignments = make(map[string]*expression)
}

func parseWire(line string) {
	tokens := identifyTokens(strings.Split(line, " "))

	/*
		fmt.Println("Line: ", line)
		for i, t := range tokens {
			fmt.Printf("  %d - %#v\n", i, t)
		}
	*/
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
	ops := []string{}
	for _, t := range operands {
		ops = append(ops, t.String())
	}
	//fmt.Printf("a:%s op:%s %d operands:%s\n", name, op.String(), len(ops), strings.Join(ops, ","))
	wireAssignments[name] = &expression{op, operands, false, 0}
}

func resolveWire(wire string) (uint16, error) {
	expr, ok := wireAssignments[wire]
	if !ok {
		return 0, fmt.Errorf("No wire named %s\n", wire)
	}
	r, err := expr.resolveExpr()
	//log.Printf("%s resolves to %d\n", wire, r)
	return r, err
}

func (expr *expression) resolveExpr() (uint16, error) {
	if expr.resolved {
		return expr.result, nil
	}
	ovalues := []uint16{}
	var value uint16
	for _, oper := range expr.operands {
		var v uint16
		switch oper.tokenType {
		case numTok:
			v = oper.value.(uint16)
		case identTok:
			v0, err := resolveWire(oper.value.(string))
			if err != nil {
				return 0, err
			}
			v = v0
		default:
			return 0, fmt.Errorf("Invalid operand %s\n", oper.String())
		}
		ovalues = append(ovalues, v)
	}
	switch expr.op.tokenType {
	case assignTok:
		value = ovalues[0]
	case notTok:
		value = ^ovalues[0]
	case andTok:
		value = ovalues[0] & ovalues[1]
	case orTok:
		value = ovalues[0] | ovalues[1]
	case lshiftTok:
		value = ovalues[0] << uint16(ovalues[1])
	case rshiftTok:
		value = ovalues[0] >> uint16(ovalues[1])
	default:
		return 0, fmt.Errorf("Invalid operator %s\n", expr.op.String())
	}
	//log.Printf("%s resolves to %d\n", expr.String(), value)
	expr.resolved = true
	expr.result = value
	return value, nil
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
	//fmt.Printf("Reading from %s\n", infile)
	reader = bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	//display := make(Display)
	for scanner.Scan() {
		line := scanner.Text()
		parseWire(line)
	}
	for _, w := range wires {
		v, err := resolveWire(w)
		if err != nil {
			log.Fatalf("Error resolving %s: %v", w, err)
		}
		fmt.Printf("%s resolves to %d\n", w, v)
	}
}
