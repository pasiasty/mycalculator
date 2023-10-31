// Code generated from f://mario//repositories//mycalculator//parser//Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcListener) ExitStart(ctx *StartContext) {}

// EnterEndLine is called when production EndLine is entered.
func (s *BaseCalcListener) EnterEndLine(ctx *EndLineContext) {}

// ExitEndLine is called when production EndLine is exited.
func (s *BaseCalcListener) ExitEndLine(ctx *EndLineContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseCalcListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseCalcListener) ExitOperation(ctx *OperationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseCalcListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseCalcListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterBracketExpression is called when production BracketExpression is entered.
func (s *BaseCalcListener) EnterBracketExpression(ctx *BracketExpressionContext) {}

// ExitBracketExpression is called when production BracketExpression is exited.
func (s *BaseCalcListener) ExitBracketExpression(ctx *BracketExpressionContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalcListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalcListener) ExitNumber(ctx *NumberContext) {}

// EnterImplicitMultipleVariableAndPower is called when production ImplicitMultipleVariableAndPower is entered.
func (s *BaseCalcListener) EnterImplicitMultipleVariableAndPower(ctx *ImplicitMultipleVariableAndPowerContext) {
}

// ExitImplicitMultipleVariableAndPower is called when production ImplicitMultipleVariableAndPower is exited.
func (s *BaseCalcListener) ExitImplicitMultipleVariableAndPower(ctx *ImplicitMultipleVariableAndPowerContext) {
}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcListener) ExitAddSub(ctx *AddSubContext) {}

// EnterRawVariableReference is called when production RawVariableReference is entered.
func (s *BaseCalcListener) EnterRawVariableReference(ctx *RawVariableReferenceContext) {}

// ExitRawVariableReference is called when production RawVariableReference is exited.
func (s *BaseCalcListener) ExitRawVariableReference(ctx *RawVariableReferenceContext) {}

// EnterImplicitVariableMultiply is called when production ImplicitVariableMultiply is entered.
func (s *BaseCalcListener) EnterImplicitVariableMultiply(ctx *ImplicitVariableMultiplyContext) {}

// ExitImplicitVariableMultiply is called when production ImplicitVariableMultiply is exited.
func (s *BaseCalcListener) ExitImplicitVariableMultiply(ctx *ImplicitVariableMultiplyContext) {}

// EnterPower is called when production Power is entered.
func (s *BaseCalcListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production Power is exited.
func (s *BaseCalcListener) ExitPower(ctx *PowerContext) {}
