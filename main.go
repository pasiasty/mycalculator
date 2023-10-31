package main

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pasiasty/mycalculator/listener"
	"github.com/pasiasty/mycalculator/parser"
)

func calc(input string) []int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := listener.New()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start_())

	return listener.Results()
}

func main() {
	lines := []string{
		"x = 3",
		"y = 2",
		"x^2",
		"2x^y",
		"2*x^2",
		"2*(x^2) + 3y",
	}
	res := calc(strings.Join(lines, "\n") + "\n")
	for idx, l := range lines {
		fmt.Printf("%s; evaluated to: %v\n", l, res[idx])
	}
}
