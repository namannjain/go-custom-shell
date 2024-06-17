package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		commands := strings.Split(input, " ")
		switch commands[0] {
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
