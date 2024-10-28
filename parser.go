package main

type Parser struct {
	lexer        *Lexer
	currentToken Token
	peekToken    Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	for p.currentToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case LET:
		return p.parseLetStatement()
	case WRITELN:
		return p.parseWritelnStatement()
	case READLN:
		return p.parseReadlnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *LetStatement {
	token := p.currentToken
	p.nextToken()
	name := &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.nextToken() // Skip the '=' token
	p.nextToken()
	value := &StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
	return &LetStatement{Token: token, Name: name, Value: value}
}

func (p *Parser) parseWritelnStatement() *WritelnStatement {
	token := p.currentToken
	p.nextToken() // Skip 'writeln'
	p.nextToken() // Skip '('
	value := &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.nextToken() // Skip identifier
	p.nextToken() // Skip ')'
	return &WritelnStatement{Token: token, Value: value}
}

func (p *Parser) parseReadlnStatement() *ReadlnStatement {
	token := p.currentToken
	p.nextToken() // Skip 'readln'
	p.nextToken() // Skip '('
	value := &Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.nextToken() // Skip identifier
	p.nextToken() // Skip ')'
	return &ReadlnStatement{Token: token, Value: value}
}
