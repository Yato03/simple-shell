package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = []string{"exit", "echo", "type"}

func execCommand(command []string) {
	if len(command) == 0 {
		return
	}

	commandName := command[0]

	switch commandName {
	case "exit":
		exitCommand(command)
	case "echo":
		echoCommand(command)
	case "type":
		typeCommand(command)
	default:
		fmt.Println(commandName + ": command not found")
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

func echoCommand(command []string) {
	message := strings.Join(command[1:], " ")
	fmt.Println(message)
}

func typeCommand(command []string) {

	if len(command) < 2 {
		fmt.Println("type: missing argument")
		return
	}

	commandName := command[1]

	if strings.Contains(strings.Join(commands, ","), commandName) {
		fmt.Println(commandName + " is a shell builtin")
	} else {
		fmt.Println(commandName + " not found")
	}

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
