package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// user_input handles user config.
//
// This function does not take any parameters.
// It does not return any values.
func user_input(wait chan struct{}, exit chan struct{}) {
	var input string
	fmt.Println("use `help` to get the list of availible commands")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		input = strings.TrimSpace(input)

		switch input {
		case "help":
			fmt.Println("help\nexit\nballs")
		case "exit":
			exit <- struct{}{}
			return
		}
	}
}
