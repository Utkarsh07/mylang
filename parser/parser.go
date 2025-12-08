package parser

import (
	"fmt"
	"mylang/ast"
	"mylang/lexer"
	"mylang/token"
)

// ---------------- Pratt's Parser ------------------

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction
)

type (
	prefixParseFunction func() ast.Expression
	infixParseFunction  func(ast.Expression) ast.Expression
)
type Parser struct {
	l                    *lexer.Lexer
	curToken             token.Token
	peekToken            token.Token
	errors               []string
	prefixParseFunctions map[token.TokenType]prefixParseFunction
	infixParseFunctions  map[token.TokenType]infixParseFunction
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) registerPrefixFunction(t token.TokenType, fn prefixParseFunction) {
	p.prefixParseFunctions[t] = fn
}

func (p *Parser) registerInfixFunction(t token.TokenType, fn infixParseFunction) {
	p.infixParseFunctions[t] = fn
}

// ----------------- Expression Parsing Functions ----------------------
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	//Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	p.prefixParseFunctions = make(map[token.TokenType]prefixParseFunction)
	p.registerPrefixFunction(token.IDENT, p.parseIdentifier)

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	message := fmt.Sprintf("Expected next token -> %s, got -> %s", t, p.peekToken.Type)
	p.errors = append(p.errors, message)
}

func (p *Parser) parseLetStatement() ast.Statement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Expression Handling
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() ast.Statement {
	statement := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Expression Handling
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFunction := p.prefixParseFunctions[p.curToken.Type]
	if prefixFunction == nil {
		return nil
	}
	leftExpression := prefixFunction()

	return leftExpression
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.curToken}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()

	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}
