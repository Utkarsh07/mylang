package lexer

import "mylang/token"

type Lexer struct {
	input        string
	position     int  // Current position in input (points to current character)
	readPosition int  // Current reading position in input (after current character)
	ch           byte // Current character under examination
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	start_position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start_position:l.position]
}

func (l *Lexer) readNumber() string {
	start_position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start_position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) makeTwoCharToken(tokenType token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	return token.Token{Type: tokenType, Literal: string(ch) + string(l.ch)}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	case '=':
		if l.peakChar() == '=' {
			t = l.makeTwoCharToken(token.EQ)
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '[':
		t = newToken(token.LBRACKET, l.ch)
	case ']':
		t = newToken(token.RBRACKET, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '!':
		if l.peakChar() == '=' {
			t = l.makeTwoCharToken(token.NOT_EQ)
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '%':
		t = newToken(token.MODULUS, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '<':
		if l.peakChar() == '=' {
			t = l.makeTwoCharToken(token.LT_EQ)
		} else {
			t = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peakChar() == '=' {
			t = l.makeTwoCharToken(token.GT_EQ)
		} else {
			t = newToken(token.GT, l.ch)
		}
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}
