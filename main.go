package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("TODO CLI started")
	fmt.Println("Type 'help' to see available commands")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the command: ")
		scan := scanner.Scan()

		if !scan {
			fmt.Println("Input error")
			break
		}

		command := scanner.Text()

		if command == "exit" {
			break
		}

		fmt.Println("You enter: ", command)
	}
}
