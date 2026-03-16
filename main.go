package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ToDO CLI started")
	fmt.Println("Type 'help' to see available command")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the command: ")
	scanner.Scan()
}
