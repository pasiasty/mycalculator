package listener

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pasiasty/mycalculator/parser"
)

type CalcListener struct {
	*parser.BaseCalcListener

	stack   []int
	results []int
	vars    map[string]int
}

func New() *CalcListener {
	return &CalcListener{
		vars: make(map[string]int),
	}
}

func (l *CalcListener) Results() []int {
	return l.results
}

func (l *CalcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *CalcListener) peek() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to peek")
	}

	return l.stack[len(l.stack)-1]
}

func (l *CalcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *CalcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *CalcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *CalcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

func (l *CalcListener) ExitEndLine(c *parser.EndLineContext) {
	l.results = append(l.results, l.pop())
}

func (l *CalcListener) ExitAssignment(c *parser.AssignmentContext) {
	l.vars[c.GetVarName().GetText()] = l.peek()
}

func (l *CalcListener) ExitImplicitVariableMultiply(c *parser.ImplicitVariableMultiplyContext) {
	multiplier, err := strconv.Atoi(c.GetVal().GetText())
	if err != nil {
		panic(err.Error())
	}
	v, ok := l.vars[c.GetVarName().GetText()]
	if !ok {
		panic(fmt.Sprintf("variable %v was undefined", c.GetText()))
	}
	l.push(multiplier * v)
}

func (l *CalcListener) ExitPower(c *parser.PowerContext) {
	right, left := l.pop(), l.pop()
	l.push(int(math.Round(math.Pow(float64(left), float64(right)))))
}

func (l *CalcListener) ExitRawVariableReference(c *parser.RawVariableReferenceContext) {
	v, ok := l.vars[c.GetText()]
	if !ok {
		panic(fmt.Sprintf("variable %v was undefined", c.GetText()))
	}
	l.push(v)
}

func (l *CalcListener) ExitImplicitMultipleVariableAndPower(c *parser.ImplicitMultipleVariableAndPowerContext) {
	multiplier, err := strconv.Atoi(c.GetMult().GetText())
	if err != nil {
		panic(err.Error())
	}
	v, ok := l.vars[c.GetVarName().GetText()]
	if !ok {
		panic(fmt.Sprintf("variable %v was undefined", c.GetText()))
	}
	exp := l.pop()
	l.push(multiplier * int(math.Round(math.Pow(float64(v), float64(exp)))))
}
