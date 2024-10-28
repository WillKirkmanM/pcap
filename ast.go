package main

import "fmt"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += s.String() + "\n"
	}
	return out
}

type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return fmt.Sprintf("%d", il.Value) }

type StringLiteral struct {
	Token Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Value }

type LetStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	return fmt.Sprintf("let %s = %s;", ls.Name.Value, ls.Value.String())
}

type PrintStatement struct {
	Token Token
	Value Expression
}

func (ps *PrintStatement) statementNode()       {}
func (ps *PrintStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PrintStatement) String() string {
	return fmt.Sprintf("print(%s);", ps.Value.String())
}

type WritelnStatement struct {
	Token Token
	Value Expression
}

func (ws *WritelnStatement) statementNode()       {}
func (ws *WritelnStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WritelnStatement) String() string {
	return fmt.Sprintf("writeln(%s);", ws.Value.String())
}

type ReadlnStatement struct {
	Token Token
	Value *Identifier
}

func (rs *ReadlnStatement) statementNode()       {}
func (rs *ReadlnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReadlnStatement) String() string {
	return fmt.Sprintf("readln(%s);", rs.Value.String())
}

type BlockStatement struct {
	Token      Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out string
	for _, s := range bs.Statements {
		out += s.String() + "\n"
	}
	return out
}
