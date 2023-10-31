package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pasiasty/mycalculator/listener"
	"github.com/pasiasty/mycalculator/parser"
)

func calc(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener listener.CalcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start_())

	return listener.Result()
}

func main() {
	cmd := "1 + 2 * 3"
	fmt.Printf("%s = %v", cmd, calc(cmd))
}
