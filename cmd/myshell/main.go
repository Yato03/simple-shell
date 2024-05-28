package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		command := input[:len(input)-1]
		// Print the input
		fmt.Println(command + ": command not found")
	}
	

}
