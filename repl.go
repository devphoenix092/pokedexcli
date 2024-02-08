package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/commands"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]

		if exists {
			var err error
			if commandName == "explore" || commandName == "catch" {
				if len(words) == 1 {
					fmt.Println("The second parameter cannot be an empty string.")
					continue
				} else {
					if commandName == "explore" {
						err = command.Callback(commands.ParamType{
							Area: words[1],
						})
					} else {
						err = command.Callback(commands.ParamType{
							PokemonName: words[1],
						})
					}
				}
			} else {
				err = command.Callback(commands.ParamType{})
			}

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
