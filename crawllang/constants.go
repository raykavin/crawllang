package crawllang

// Bytecode instructions
const (
	OP_NAVIGATE = iota
	OP_CLICK
	OP_STORE_VAR
	OP_LOAD_VAR
	OP_PUSH_CONST
	OP_HALT
)

const (
	TOKEN_EOF = iota
	TOKEN_ILLEGAL

	// Functions
	TOKEN_NAVIGATE
	TOKEN_CLICK

	TOKEN_NUMBER
	TOKEN_STRING

	// Delimiters
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_SEMICOLON

	// Operators
	TOKEN_ASSIGN
	TOKEN_PLUS

	// Keywords
	TOKEN_VAR
	TOKEN_FUNCTION
)

const (
	NAVIGATE_FUNC string = "NAVIGATE"
	CLICK_FUNC    string = "CLICK"

	LPAREN     byte = '('
	RPAREN     byte = ')'
	LBRACE     byte = '{'
	RBRACE     byte = '}'
	SEMICOLON  byte = ';'
	ASSIGNMENT byte = '='
	STRING     byte = '"'
)
