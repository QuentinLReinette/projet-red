package main

import (
	"fmt"
	"strconv"
	"strings"
)

type menuOption struct {
	text   string
	action func()
}

func menuPrint(options []menuOption) {
	for {
		println()
		for i, o := range options {
			fmt.Printf(" %d. %s\n", i+1, o.text)
		}
		println(" 0. Back")

		Scanner.Scan()

		input := strings.TrimSpace(Scanner.Text())
		if input == "" {
			continue
		}
		choice, err := strconv.Atoi(input)
		if err != nil {
			println("Please enter a number")
		}
		if choice < 0 || choice > len(options) {
			println("Invalid choice")
			continue
		}
		if choice == 0 {
			return
		}
		options[choice-1].action()

		if err := Scanner.Err(); err != nil {
			fmt.Printf("Error while reading input: %s", err)
		}
	}
}
