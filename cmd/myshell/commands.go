package main

import (
	"fmt"
	"log"
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
		if value.getName() == commandName {
			fmt.Println(commandName + " is a shell builtin")
			return
		}
	}

	//search in PATH
	dir, ok := searchCommandInPath(commandName)
	if ok {
		fmt.Println(commandName + " is " + dir)
		return
	}

	fmt.Println(commandName + " not found")

}

func (c *TypeCommand) getName() string {
	return c.name
}

type PwdCommand struct {
	name string
}

/*
func (c *PwdCommand) execute(args []string) {
	path, ok := os.LookupEnv("PWD")
	if ok {
		fmt.Println(path)
	}
}*/

func (c *PwdCommand) execute(args []string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}

func (c *PwdCommand) getName() string {
	return c.name
}

type ChangeDirectoryCommand struct {
	name string
}

func (c *ChangeDirectoryCommand) execute(args []string) {

	if len(args) < 1 {
		fmt.Println("cd: missing argument")
		return
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println("cd: cannot access the current path")
		return
	}
	err = os.Chdir(filepath.Join(path, args[0]))
	if err != nil {
		fmt.Println("cd: " + path + ": No such file or directory")
	}
}

func (c *ChangeDirectoryCommand) getName() string {
	return c.name
}

var commands = []Command{
	&ExitCommand{name: "exit"},
	&EchoCommand{name: "echo"},
	&TypeCommand{name: "type"},
	&PwdCommand{name: "pwd"},
	&ChangeDirectoryCommand{name: "cd"},
}
