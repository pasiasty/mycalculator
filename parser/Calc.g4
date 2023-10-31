grammar Calc;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
VARIABLE_NAME: [a-z][a-z0-9]*;
OPEN_BRACKET: '(';
CLOSE_BRACKET: ')';
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: line (line)* EOF;

line: operation '\n' # EndLine;

operation: expression | assignment;

assignment: VarName = VARIABLE_NAME '=' expression;

expression:
	'(' expression ')'										# BracketExpression
	| mult = NUMBER VarName = VARIABLE_NAME '^' expression	# ImplicitMultipleVariableAndPower
	| expression '^' expression								# Power
	| expression op = ('*' | '/') expression				# MulDiv
	| expression op = ('+' | '-') expression				# AddSub
	| NUMBER												# Number
	| VARIABLE_NAME											# RawVariableReference
	| val = NUMBER VarName = VARIABLE_NAME					# ImplicitVariableMultiply;