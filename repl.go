package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MarunDArbaumont/pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
			fmt.Print("Pokedex > ")
			scanner.Scan()

			words := cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}

			firstWord := words[0]
			command, exists := getCommands()[firstWord]

			if exists {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknon command")
				continue
			}
        }
}

func cleanInput(text string) []string {
        text = strings.ToLower(text)
        return strings.Fields(text)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world",
			callback: commandMapb,
		},
	}
}