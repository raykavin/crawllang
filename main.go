package main

import (
	"crawlang/crawllang"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "crawllang",
		Short: "crawllang - Lightweight interpreted language for web scraping",
	}

	rootCmd.AddCommand(newCompileCmd())
	rootCmd.AddCommand(newRunCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func newCompileCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "compile <file>",
		Short: "Compile source code to bytecode",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]

			content, err := os.ReadFile(filename)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}
			sourceCode := string(content)

			compiler := crawllang.NewCompiler()
			bytecode := compiler.Compile(sourceCode)

			if err := crawllang.SaveBytecode(output, bytecode); err != nil {
				return fmt.Errorf("failed to save bytecode: %w", err)
			}

			fmt.Printf("Successfully compiled to: %s\n", output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file path (required)")
	cmd.MarkFlagRequired("output")

	return cmd
}

func newRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run <bytecode-file>",
		Short: "Execute compiled bytecode",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			bytecode, err := crawllang.LoadBytecode(args[0])
			if err != nil {
				return fmt.Errorf("failed to load bytecode: %w", err)
			}

			crawllang.NewVM(bytecode).Run()
			return nil
		},
	}
}
