package crawllang

import "fmt"

// Interpreter executes parsed commands
type Interpreter struct {
	commands  map[string]func(string)
	variables map[string]string
}

// NewInterpreter creates a new interpreter with command mappings
func NewInterpreter() *Interpreter {
	i := &Interpreter{
		commands:  make(map[string]func(string)),
		variables: make(map[string]string),
	}

	i.commands["NAVIGATE"] = func(param string) {
		fmt.Printf("Navigating to: %s\n", param)
	}

	i.commands["CLICK"] = func(param string) {
		fmt.Printf("Clicking element: %s\n", param)
	}

	return i
}

// Declares and assigns a variable
func (i *Interpreter) declareVariable(name, value string) {
	i.variables[name] = value
}

// Retrieves a variable's value
func (i *Interpreter) useVariable(name string) string {
	return i.variables[name]
}

// Executes a command with parameter
func (i *Interpreter) executeCommand(command, param string) {
	if action, exists := i.commands[command]; exists {
		action(param)
	}
}
