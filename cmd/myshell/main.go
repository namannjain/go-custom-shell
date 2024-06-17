package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the command: ", err)
	}
	command = strings.Trim(command, "\n")
	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}
