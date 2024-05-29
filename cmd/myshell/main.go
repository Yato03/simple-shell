package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func execCommand(command []string) {
	switch command[0] {
	case "exit":
		exitCommand(command)
	default:
		fmt.Println(command[0] + ": command not found")
	}
}

func exitCommand(command []string) {
	if len(command) < 2 {
		os.Exit(0)
	}
	exitCode, err := strconv.Atoi(command[1])
	if err != nil {
		os.Exit(0)
	}
	os.Exit(exitCode)
}

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		command := input[:len(input)-1]

		// Split the input into a slice of strings
		commandSlice := strings.Split(command, " ")

		// Execute the command
		execCommand(commandSlice)
	}

}
