package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	builtinCommands := []string{"echo", "exit", "cat", "type"}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		commands := strings.Split(input, " ")

		switch commands[0] {
		case "type":
			if slices.Contains(builtinCommands, commands[1]) {
				fmt.Printf("%s is a shell builtin\n", commands[1])
			} else {
				fmt.Printf("%s: not found\n", commands[1])
			}

		case "exit":
			exitCode, err := strconv.Atoi(commands[1])
			if err == nil {
				os.Exit(exitCode)
			}

		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(commands[1:], " "))

		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
