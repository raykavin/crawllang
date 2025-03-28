package crawllang

import (
	"encoding/gob"
	"os"
)

func SaveBytecode(filename string, bytecode []Instruction) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(bytecode)
	if err != nil {
		return err
	}

	return nil
}

func LoadBytecode(filename string) ([]Instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bytecode []Instruction
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&bytecode)
	if err != nil {
		return nil, err
	}

	return bytecode, nil
}
