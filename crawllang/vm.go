package crawllang

import "fmt"

// Virtual Machine implementation
type VM struct {
	stack     []string
	variables map[string]string
	bytecode  []Instruction
	pc        int
}

func NewVM(bytecode []Instruction) *VM {
	return &VM{
		variables: make(map[string]string),
		bytecode:  bytecode,
	}
}

func (v *VM) Run() {
	for v.pc < len(v.bytecode) {
		instr := v.bytecode[v.pc]
		v.pc++

		switch instr.Opcode {
		case OP_PUSH_CONST:
			v.stack = append(v.stack, instr.Operand)

		case OP_STORE_VAR:
			if len(v.stack) == 0 {
				panic("stack underflow")
			}
			val := v.stack[len(v.stack)-1]
			v.stack = v.stack[:len(v.stack)-1]
			v.variables[instr.Operand] = val

		case OP_LOAD_VAR:
			val, exists := v.variables[instr.Operand]
			if !exists {
				panic("undefined variable: " + instr.Operand)
			}
			v.stack = append(v.stack, val)

		case OP_NAVIGATE:
			if len(v.stack) == 0 {
				panic("stack underflow")
			}
			url := v.stack[len(v.stack)-1]
			v.stack = v.stack[:len(v.stack)-1]
			fmt.Printf("Navigating to: %s\n", url)

		case OP_CLICK:
			if len(v.stack) == 0 {
				panic("stack underflow")
			}
			element := v.stack[len(v.stack)-1]
			v.stack = v.stack[:len(v.stack)-1]
			fmt.Printf("Clicking element: %s\n", element)

		case OP_HALT:
			return
		}
	}
}
