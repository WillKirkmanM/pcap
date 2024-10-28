package main

import (
	"log"
	"strings"
)

func generateCode(program *Program) string {
	var out strings.Builder

	for _, stmt := range program.Statements {
		if stmt == nil {
			log.Println("generateCode: encountered nil statement")
			continue
		}
		switch stmt := stmt.(type) {
		case *LetStatement:
			out.WriteString(generateLetStatement(stmt))
		case *WritelnStatement:
			out.WriteString(generateWritelnStatement(stmt))
		case *ReadlnStatement:
			out.WriteString(generateReadlnStatement(stmt))
		default:
			log.Printf("generateCode: encountered unknown statement type %T", stmt)
		}
	}

	return out.String()
}

func generateLetStatement(stmt *LetStatement) string {
	if stmt == nil || stmt.Name == nil || stmt.Value == nil {
		log.Println("generateLetStatement: encountered nil fields")
		return ""
	}
	switch value := stmt.Value.(type) {
	case *StringLiteral:
		return "MOV " + stmt.Name.Value + ", '" + value.Value + "'\n"
	case *Identifier:
		return "MOV " + stmt.Name.Value + ", " + value.Value + "\n"
	default:
		log.Printf("generateLetStatement: encountered unknown value type %T", value)
		return ""
	}
}

func generateWritelnStatement(stmt *WritelnStatement) string {
	if stmt == nil || stmt.Value == nil {
		log.Println("generateWritelnStatement: encountered nil fields")
		return ""
	}

	switch value := stmt.Value.(type) {
	case *Identifier:
		return "WRITELN " + value.Value + "\n"
	case *StringLiteral:
		return "WRITELN '" + value.Value + "'\n"
	default:
		log.Printf("generateWritelnStatement: encountered unknown value type %T", value)
		return ""
	}
}

func generateReadlnStatement(stmt *ReadlnStatement) string {
	if stmt == nil || stmt.Value == nil {
		log.Println("generateReadlnStatement: encountered nil fields")
		return ""
	}

	return "READLN " + stmt.Value.Value + "\n"
}
