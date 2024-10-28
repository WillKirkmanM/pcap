package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Interpreter struct {
	variables map[string]string
}

func NewInterpreter() *Interpreter {
	return &Interpreter{variables: make(map[string]string)}
}

func (i *Interpreter) Interpret(code string) {
	lines := strings.Split(code, "\n")
	reader := bufio.NewReader(os.Stdin)
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Printf("Invalid line: %s\n", line)
			continue
		}
		command := parts[0]
		switch command {
		case "MOV":
			if len(parts) < 3 {
				fmt.Printf("Invalid MOV command: %s\n", line)
				continue
			}
			variable := strings.Trim(parts[1], ",")
			value := strings.Join(parts[2:], " ")
			value = strings.Trim(value, "'")
			i.variables[variable] = value
		case "WRITELN":
			if len(parts) != 2 {
				fmt.Printf("Invalid WRITELN command: %s\n", line)
				continue
			}
			variable := parts[1]
			value, ok := i.variables[variable]
			if !ok {
				fmt.Printf("Undefined variable: %s\n", variable)
				continue
			}
			fmt.Println(value)
		case "READLN":
			if len(parts) != 2 {
				fmt.Printf("Invalid READLN command: %s\n", line)
				continue
			}
			variable := parts[1]
			fmt.Printf("Enter value for %s: ", variable)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				continue
			}
			input = strings.TrimSpace(input)
			i.variables[variable] = input
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}
