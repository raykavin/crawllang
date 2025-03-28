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
	l.skipWhitespace()

	// Map of known single-character token handlers
	charHandlers := map[byte]func() Token{
		LPAREN: func() Token {
			tok := Token{Type: TOKEN_LPAREN, Literal: string(l.ch)}
			l.readChar()
			return tok
		},
		RPAREN: func() Token {
			tok := Token{Type: TOKEN_RPAREN, Literal: string(l.ch)}
			l.readChar()
			return tok
		},
		SEMICOLON: func() Token {
			tok := Token{Type: TOKEN_SEMICOLON, Literal: string(l.ch)}
			l.readChar()
			return tok
		},
		ASSIGNMENT: func() Token {
			tok := Token{Type: TOKEN_ASSIGN, Literal: string(l.ch)}
			l.readChar()
			return tok
		},
		STRING: func() Token {
			lit := l.readString()
			return Token{Type: TOKEN_STRING, Literal: lit}
		},
	}

	// Check if we have a handler for the current character
	if handler, exists := charHandlers[l.ch]; exists {
		return handler()
	}

	// Handle special cases
	switch {
	case l.ch == 0:
		tok := Token{Type: TOKEN_EOF}
		l.readChar()
		return tok

	case isLetter(l.ch):
		literal := l.readIdentifier()
		tokType := TOKEN_VAR

		switch literal {
		case NAVIGATE_FUNC:
			tokType = TOKEN_NAVIGATE
		case CLICK_FUNC:
			tokType = TOKEN_CLICK
		}

		return Token{Type: tokType, Literal: literal}

	case isDigit(l.ch):
		literal := l.readNumber()
		return Token{Type: TOKEN_NUMBER, Literal: literal}
	}

	// Handle unknown characters
	tok := Token{Type: TOKEN_ILLEGAL, Literal: string(l.ch)}
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
	// Skip opening quote
	l.readChar()

	start := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	end := l.position

	// Skip closing quote
	l.readChar()
	return l.input[start:end]
}
