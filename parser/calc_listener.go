// Code generated from f://mario//repositories//mycalculator//parser//Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterEndLine is called when entering the EndLine production.
	EnterEndLine(c *EndLineContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterBracketExpression is called when entering the BracketExpression production.
	EnterBracketExpression(c *BracketExpressionContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterImplicitMultipleVariableAndPower is called when entering the ImplicitMultipleVariableAndPower production.
	EnterImplicitMultipleVariableAndPower(c *ImplicitMultipleVariableAndPowerContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterRawVariableReference is called when entering the RawVariableReference production.
	EnterRawVariableReference(c *RawVariableReferenceContext)

	// EnterImplicitVariableMultiply is called when entering the ImplicitVariableMultiply production.
	EnterImplicitVariableMultiply(c *ImplicitVariableMultiplyContext)

	// EnterPower is called when entering the Power production.
	EnterPower(c *PowerContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitEndLine is called when exiting the EndLine production.
	ExitEndLine(c *EndLineContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitBracketExpression is called when exiting the BracketExpression production.
	ExitBracketExpression(c *BracketExpressionContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitImplicitMultipleVariableAndPower is called when exiting the ImplicitMultipleVariableAndPower production.
	ExitImplicitMultipleVariableAndPower(c *ImplicitMultipleVariableAndPowerContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitRawVariableReference is called when exiting the RawVariableReference production.
	ExitRawVariableReference(c *RawVariableReferenceContext)

	// ExitImplicitVariableMultiply is called when exiting the ImplicitVariableMultiply production.
	ExitImplicitVariableMultiply(c *ImplicitVariableMultiplyContext)

	// ExitPower is called when exiting the Power production.
	ExitPower(c *PowerContext)
}
