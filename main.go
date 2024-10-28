package main

import "os"

func main() {
	content, _ := os.ReadFile("example.pas")

	input := string(content)

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()
	code := generateCode(program)
	interpreter := NewInterpreter()
	interpreter.Interpret(code)
}
