package crawllang

// Parser builds AST from tokens
type Parser struct {
	lexer  *Lexer
	curTok Token
}

// NewParser creates a new parser instance
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

// Advances to the next token
func (p *Parser) nextToken() {
	p.curTok = p.lexer.nextToken()
}
