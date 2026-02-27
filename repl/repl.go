package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	supportedCommands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name : "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
	for ;; {
			fmt.Print("Pokedex > ")
			scanner.Scan()

			words := cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}

			firstWord := words[0]
			command, exists := supportedCommands[firstWord]

			if exists {
				command.callback()
			} else {
				fmt.Print("Unknon command")
			}
        }
}

func cleanInput(text string) []string {
        text = strings.ToLower(text)
        return strings.Fields(text)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}