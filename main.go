package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	char1 *character
)

func main() {
	char1 = initChar("newCharacter", "elf", 100)
	printMenu()
}

func printMenu() {
	fmt.Println("\n 1. Show character's info\n" +
		" 2. Show inventory\n" +
		" 3. Exit")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		choice, err := strconv.Atoi(input)
		if err != nil {
			println("Please enter a number")
			continue
		}
		switch choice {
		case 1:
			char1.displayInfo()
			printMenu()
			return
		case 2:
			// inventory
		case 3:
			os.Exit(0)
		default:
			println("Invalid choice")
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error while reading input: ", err)
		}
	}
}
