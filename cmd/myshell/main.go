package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func execCommand(command []string) {
	if len(command) == 0 {
		return
	}

	commandName := command[0]
	args := command[1:]

	//search builint commands
	for _, c := range commands {
		if c.getName() == commandName {
			c.execute(args)
			return
		}
	}

	//search in PATH
	path, ok := searchCommandInPath(commandName)
	if ok {
		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println(commandName + ": command not found")

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
