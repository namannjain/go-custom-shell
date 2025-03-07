package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func isPresent(arr []string, target string) bool {
	found := false
	for _, ele := range arr {
		if ele == target {
			found = true
			break
		}
	}
	return found
}

func getFilePathIfExists(directories []string, fileName string) (string, error) {
	for _, directory := range directories {
		filePath := directory + "/" + fileName
		_, err := os.Stat(filePath)
		if err == nil {
			return filePath, nil
		}
	}

	return "", errors.New("file does not exist")
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//env path
	directories := strings.Split(os.Getenv("PATH"), ":")

	reader := bufio.NewReader(os.Stdin)
	builtinCommands := []string{"echo", "exit", "type", "pwd", "cd"}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		commands := strings.Split(input, " ")

		switch commands[0] {

		case "type":
			if isPresent(builtinCommands, commands[1]) {
				fmt.Printf("%s is a shell builtin\n", commands[1])
			} else if filepath, err := getFilePathIfExists(directories, commands[1]); err == nil {
				fmt.Printf("%s is %s\n", commands[1], filepath)
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

		case "pwd":
			if currwd, err := os.Getwd(); err == nil {
				fmt.Println(currwd)
			}

		case "cd":
			if commands[1] == "~" {
				commands[1] = os.Getenv("HOME")
			}

			resp, err := exists(commands[1])
			if resp && err == nil {
				os.Chdir(commands[1])
			} else {
				fmt.Printf("cd: %s: No such file or directory\n", commands[1])
			}

		default:
			_, err := getFilePathIfExists(directories, commands[0])
			if err != nil {
				fmt.Printf("%s: command not found\n", input)
			} else {
				exernalProgramObj := exec.Command(commands[0])
				arguments := commands[1:]
				exernalProgramObj.Args = append(exernalProgramObj.Args, arguments...)

				// externalProgramOutput, err := exernalProgramObj.CombinedOutput()
				exernalProgramObj.Stderr = os.Stderr
				exernalProgramObj.Stdout = os.Stdout
				exernalProgramObj.Run()
			}
		}
	}
}
