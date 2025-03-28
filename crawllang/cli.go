package crawllang

import "fmt"

func CliInterpreter(code string) {
	lexer := NewLexer(code)
	parser := NewParser(lexer)
	interpreter := NewInterpreter()

	for parser.curTok.Type != TOKEN_EOF {
		switch {
		case parser.curTok.Type == TOKEN_VAR:
			parser.nextToken() // Consume 'var'
			varName := parser.curTok.Literal
			parser.nextToken() // Consume variable name

			if parser.curTok.Type != TOKEN_ASSIGN {
				fmt.Println("Expected assignment")
				break
			}

			parser.nextToken() // Consume '='
			var varValue string

			switch parser.curTok.Type {
			case TOKEN_VAR:
				varValue = interpreter.useVariable(parser.curTok.Literal)
			case TOKEN_STRING, TOKEN_NUMBER:
				varValue = parser.curTok.Literal
			default:
				fmt.Println("Invalid value type")
			}

			interpreter.declareVariable(varName, varValue)
			parser.nextToken() // Consume value
			if parser.curTok.Type == TOKEN_SEMICOLON {
				parser.nextToken() // Consume ';'
			}

		case parser.curTok.Type == TOKEN_NAVIGATE:
			parser.nextToken() // Consume 'NAVIGATE'

			if parser.curTok.Type != TOKEN_LPAREN {
				fmt.Println("Expected '('")
				break
			}

			parser.nextToken() // Consume '('
			var param string

			switch parser.curTok.Type {
			case TOKEN_VAR:
				param = interpreter.useVariable(parser.curTok.Literal)
			case TOKEN_STRING, TOKEN_NUMBER:
				param = parser.curTok.Literal
			default:
				fmt.Println("Invalid parameter type")
			}

			parser.nextToken() // Consume parameter
			if parser.curTok.Type != TOKEN_RPAREN {
				fmt.Println("Expected ')'")
				break
			}

			parser.nextToken() // Consume ')'
			if parser.curTok.Type == TOKEN_SEMICOLON {
				interpreter.executeCommand("NAVIGATE", param)
				parser.nextToken() // Consume ';'
			}

		case parser.curTok.Type == TOKEN_CLICK:
			parser.nextToken() // Consume 'click'

			if parser.curTok.Type != TOKEN_LPAREN {
				fmt.Println("Expected '('")
				break
			}

			parser.nextToken() // Consume '('
			var param string

			switch parser.curTok.Type {
			case TOKEN_VAR:
				param = interpreter.useVariable(parser.curTok.Literal)
			case TOKEN_STRING, TOKEN_NUMBER:
				param = parser.curTok.Literal
			default:
				fmt.Println("Invalid parameter type")
			}

			parser.nextToken() // Consume parameter
			if parser.curTok.Type != TOKEN_RPAREN {
				fmt.Println("Expected ')'")
				break
			}

			parser.nextToken() // Consume ')'
			if parser.curTok.Type == TOKEN_SEMICOLON {
				interpreter.executeCommand("CLICK", param)
				parser.nextToken() // Consume ';'
			}

		default:
			parser.nextToken()
		}
	}
}
