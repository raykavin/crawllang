package crawllang

// Compiler implementation
type Instruction struct {
	Opcode  int
	Operand string
}

type Compiler struct {
	bytecode  []Instruction
	variables map[string]int
}

func NewCompiler() *Compiler {
	return &Compiler{
		variables: make(map[string]int),
	}
}

func (c *Compiler) Compile(input string) []Instruction {
	lexer := NewLexer(input)
	parser := NewParser(lexer)

	for parser.curTok.Type != TOKEN_EOF {
		switch {
		case parser.curTok.Type == TOKEN_VAR:
			c.compileVarDeclaration(parser)
		case parser.curTok.Type == TOKEN_NAVIGATE:
			c.compileNavigate(parser)
		case parser.curTok.Type == TOKEN_CLICK:
			c.compileClick(parser)
		default:
			parser.nextToken()
		}
	}

	c.bytecode = append(c.bytecode, Instruction{Opcode: OP_HALT})
	return c.bytecode
}

// func (c *Compiler) Serialize() string {
// 	var sb strings.Builder
// 	for _, instr := range c.bytecode {
// 		sb.WriteString(fmt.Sprintf("%d|%s\n", instr.Opcode, instr.Operand))
// 	}
// 	return sb.String()
// }

// func Deserialize(data string) []Instruction {
// 	var bytecode []Instruction
// 	lines := strings.Split(data, "\n")
// 	for _, line := range lines {
// 		if line == "" {
// 			continue
// 		}
// 		parts := strings.SplitN(line, "|", 2)
// 		if len(parts) != 2 {
// 			continue
// 		}
// 		var opcode int
// 		fmt.Sscanf(parts[0], "%d", &opcode)
// 		operand := parts[1]
// 		bytecode = append(bytecode, Instruction{Opcode: opcode, Operand: operand})
// 	}
// 	return bytecode
// }

// func (c *Compiler) WriteBytecode(output string) error {
// 	// var builder strings.Builder

// 	// for _, instr := range c.bytecode {
// 	// 	name, exists := opcodeNames[instr.Opcode]
// 	// 	if !exists {
// 	// 		name = "UNKNOWN"
// 	// 	}

// 	// 	operand := ""
// 	// 	if instr.Operand != "" {
// 	// 		// Show quoted strings for better readability
// 	// 		if strings.Contains(instr.Operand, `"`) {
// 	// 			operand = fmt.Sprintf("%q", instr.Operand)
// 	// 		} else {
// 	// 			operand = instr.Operand
// 	// 		}
// 	// 	}

// 	// 	if operand != "" {
// 	// 		builder.Write(fmt.Appendf(nil, "%-12s %s\n", name+":", operand))
// 	// 	} else {
// 	// 		builder.Write([]byte(name + ":"))
// 	// 	}
// 	// }

// 	return os.WriteFile(output, []byte(c.Serialize()), os.ModePerm)
// }

func (c *Compiler) compileVarDeclaration(p *Parser) {
	p.nextToken() // Consume VAR
	varName := p.curTok.Literal
	p.nextToken() // Consume variable name
	p.nextToken() // Consume ASSIGN

	switch p.curTok.Type {
	case TOKEN_STRING, TOKEN_NUMBER:
		c.bytecode = append(c.bytecode, Instruction{
			Opcode:  OP_PUSH_CONST,
			Operand: p.curTok.Literal,
		})
	case TOKEN_VAR:
		c.bytecode = append(c.bytecode, Instruction{
			Opcode:  OP_LOAD_VAR,
			Operand: p.curTok.Literal,
		})
	}

	c.bytecode = append(c.bytecode, Instruction{
		Opcode:  OP_STORE_VAR,
		Operand: varName,
	})

	p.nextToken() // Consume value
	p.nextToken() // Consume SEMICOLON
}

func (c *Compiler) compileNavigate(p *Parser) {
	p.nextToken() // Consume NAVIGATE
	p.nextToken() // Consume LPAREN

	switch p.curTok.Type {
	case TOKEN_VAR:
		c.bytecode = append(c.bytecode, Instruction{
			Opcode:  OP_LOAD_VAR,
			Operand: p.curTok.Literal,
		})
	case TOKEN_STRING, TOKEN_NUMBER:
		c.bytecode = append(c.bytecode, Instruction{
			Opcode:  OP_PUSH_CONST,
			Operand: p.curTok.Literal,
		})
	}

	c.bytecode = append(c.bytecode, Instruction{Opcode: OP_NAVIGATE})
	p.nextToken() // Consume param
	p.nextToken() // Consume RPAREN
	p.nextToken() // Consume SEMICOLON
}

func (c *Compiler) compileClick(p *Parser) {
	p.nextToken() // Consume CLICK
	p.nextToken() // Consume LPAREN

	switch p.curTok.Type {
	case TOKEN_VAR:
		c.bytecode = append(c.bytecode, Instruction{
			Opcode:  OP_LOAD_VAR,
			Operand: p.curTok.Literal,
		})
	case TOKEN_STRING, TOKEN_NUMBER:
		c.bytecode = append(c.bytecode, Instruction{
			Opcode:  OP_PUSH_CONST,
			Operand: p.curTok.Literal,
		})
	}

	c.bytecode = append(c.bytecode, Instruction{Opcode: OP_CLICK})
	p.nextToken() // Consume param
	p.nextToken() // Consume RPAREN
	p.nextToken() // Consume SEMICOLON
}
