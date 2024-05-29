package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Command interface {
	execute(args []string)
	getName() string
}

type ExitCommand struct {
	name string
}

func (c *ExitCommand) execute(args []string) {
	if len(args) < 1 {
		os.Exit(0)
	}
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		os.Exit(0)
	}
	os.Exit(exitCode)
}

func (c *ExitCommand) getName() string {
	return c.name
}

type EchoCommand struct {
	name string
}

func (c *EchoCommand) execute(args []string) {
	message := strings.Join(args, " ")
	fmt.Println(message)
}

func (c *EchoCommand) getName() string {
	return c.name
}

type TypeCommand struct {
	name string
}

func (c *TypeCommand) execute(args []string) {
	if len(args) < 1 {
		fmt.Println("type: missing argument")
		return
	}

	commandName := args[0]

	//search builint commands
	for _, value := range commands {
		if value == commandName {
			fmt.Println(commandName + " is a shell builtin")
			return
		}
	}

	//search in PATH
	path, ok := os.LookupEnv("PATH")

	if ok {
		dirs := strings.Split(path, ":")
		for _, dir := range dirs {
			if listFiles(dir, commandName) {
				fmt.Println(commandName + " is " + dir + "/" + commandName)
				return
			}
		}
	}

	fmt.Println(commandName + " not found")

}

func (c *TypeCommand) getName() string {
	return c.name
}

var commands = []string{"exit", "echo", "type"}
var commands2 = []Command{
	&ExitCommand{name: "exit"},
	&EchoCommand{name: "echo"},
	&TypeCommand{name: "type"},
}

func listFiles(dir string, name string) bool {
	result := false
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, "/"+name) {
			//fmt.Println(path)
			result = true
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
	}
	return result
}

func execCommand(command []string) {
	if len(command) == 0 {
		return
	}

	commandName := command[0]
	args := command[1:]

	for _, c := range commands2 {
		if c.getName() == commandName {
			c.execute(args)
			return
		}
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
