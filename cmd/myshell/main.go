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
		bufio.NewReader(os.Stdin).ReadString('\n')

		// Print the input
		fmt.Println(input + ": command not found")
	}
	

}
