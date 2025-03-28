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
	opHandlers := map[int]func(*VM, string) bool{
		OP_PUSH_CONST: (*VM).handlePushConst,
		OP_STORE_VAR:  (*VM).handleStoreVar,
		OP_LOAD_VAR:   (*VM).handleLoadVar,
		OP_NAVIGATE:   (*VM).handleNavigate,
		OP_CLICK:      (*VM).handleClick,
		OP_HALT:       (*VM).handleHalt,
	}

	for v.pc < len(v.bytecode) {
		instr := v.bytecode[v.pc]
		v.pc++

		handler, ok := opHandlers[instr.Opcode]
		if !ok {
			panic(fmt.Sprintf("unknown opcode: %v", instr.Opcode))
		}

		if !handler(v, instr.Operand) {
			return
		}
	}
}

func (v *VM) handlePushConst(operand string) bool {
	v.stack = append(v.stack, operand)
	return true
}

func (v *VM) handleStoreVar(operand string) bool {
	if len(v.stack) == 0 {
		panic("stack underflow")
	}
	val := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	v.variables[operand] = val
	return true
}

func (v *VM) handleLoadVar(operand string) bool {
	val, exists := v.variables[operand]
	if !exists {
		panic("undefined variable: " + operand)
	}
	v.stack = append(v.stack, val)
	return true
}

func (v *VM) handleNavigate(operand string) bool {
	if len(v.stack) == 0 {
		panic("stack underflow")
	}
	url := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	fmt.Printf("Navigating to: %s\n", url)
	return true
}

func (v *VM) handleClick(operand string) bool {
	if len(v.stack) == 0 {
		panic("stack underflow")
	}
	element := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	fmt.Printf("Clicking element: %s\n", element)
	return true
}

func (v *VM) handleHalt(operand string) bool {
	return false
}
