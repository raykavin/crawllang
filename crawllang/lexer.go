package crawllang

// Lexer converts input text into tokens
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// NewLexer creates a new lexer instance
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Reads the next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Reads the next token
func (l *Lexer) nextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = Token{Type: TOKEN_LPAREN, Literal: string(l.ch)}
	case ')':
		tok = Token{Type: TOKEN_RPAREN, Literal: string(l.ch)}
	case ';':
		tok = Token{Type: TOKEN_SEMICOLON, Literal: string(l.ch)}
	case '=':
		tok = Token{Type: TOKEN_ASSIGN, Literal: string(l.ch)}
	case '"':
		tok.Type = TOKEN_STRING
		tok.Literal = l.readString()
		return tok
	case 0:
		tok = Token{Type: TOKEN_EOF}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			switch tok.Literal {
			case "NAVIGATE":
				tok.Type = TOKEN_NAVIGATE
			case "CLICK":
				tok.Type = TOKEN_CLICK
			default:
				tok.Type = TOKEN_VAR
			}
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = TOKEN_NUMBER
			return tok
		}
	}

	l.readChar()
	return tok
}

// Skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Reads an identifier
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// Reads a number literal
func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// Reads a string literal
func (l *Lexer) readString() string {
	l.readChar() // Skip opening quote
	start := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	end := l.position
	l.readChar() // Skip closing quote
	return l.input[start:end]
}
