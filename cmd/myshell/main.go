package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = map[string]func([]string){
	"exit": exitCommand,
	"echo": echoCommand,
}

func execCommand(command []string) {
	if len(command) == 0 {
		return
	}

	function, ok := commands[command[0]]
	// If the key exists
	if ok {
		function(command)
	} else {
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

func echoCommand(command []string) {
	message := strings.Join(command[1:], " ")
	fmt.Println(message)
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
