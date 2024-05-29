package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

	for _, c := range commands {
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
