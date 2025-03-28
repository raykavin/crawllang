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

// var opcodeNames = map[int]string{
// 	OP_NAVIGATE:   "NAVIGATE",
// 	OP_CLICK:      "CLICK",
// 	OP_STORE_VAR:  "STORE_VAR",
// 	OP_LOAD_VAR:   "LOAD_VAR",
// 	OP_PUSH_CONST: "PUSH_CONST",
// 	OP_HALT:       "HALT",
// }

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
