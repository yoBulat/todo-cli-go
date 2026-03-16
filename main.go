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

	Loop:
	for {
		fmt.Print("Enter the command: ")
		scan := scanner.Scan()

		if !scan {
			fmt.Println("Input error")
			break
		}

		command := scanner.Text()

		switch command {
		case "help":
			fmt.Println(`Available commands:
			-add  -Adds a task
			-list  -Show all tasks
			-del  -Delete a task
			-exit -Exit the program`)
		case "exit":
			break Loop
		case "":
			continue
		default:
			fmt.Println("The command could not be recognized")
		}
	}
}
