# CrawlLang

CrawlLang is a lightweight interpreted language designed for web scraping. It provides a simple and efficient way to compile and execute bytecode from source code files.

## Features

- Compile source code to bytecode.
- Execute compiled bytecode using a virtual machine.
- Built with Go and uses the Cobra library for CLI commands.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/raykavin/crawllang.git
   cd crawllang
   go build -o crawllang 
   ```

## Usage

CrawlLang supports two main commands: compile and run.

## Compile

Compiles a source code file to bytecode.
   
   ```bash
   ./crawllang compile <file> -o <output>
   ```

## Arguments

- `<file>`: Path to the source code file.
- `-o`, `--output`: Output file path for the compiled bytecode (required).

## Example

```bash
./crawllang compile example.cwl -o example.cwb
```

## Run

Executes a compiled bytecode file.

```bash
./crawllang run <bytecode-file>

```
## Arguments

- `<bytecode-file>`: Path to the compiled bytecode file.

## Example

```bash
./crawllang run example.cwb
```


# Contributing

Feel free to open issues or pull requests to contribute to this project. Your feedback and improvements are welcome!
