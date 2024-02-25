package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands_builder "github.com/nursakn/pokedexcli/commands"
	"github.com/nursakn/pokedexcli/config"
)

func startRepl(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := commands_builder.GetCommands()

	fmt.Print("Welcome to Pokedex Cli\n> ")
	for scanner.Scan() {
		commandString := scanner.Text()
		words := strings.Fields(commandString)

		if len(words) == 0 {
			fmt.Print("> ")
			continue
		}

		err := commands.Send(cleanInput(words[0]), cfg, words[1:])

		if err != nil {
			errorString := fmt.Sprintf("error: %s", err.Error())
			fmt.Println(errorString)
		}

		fmt.Print("> ")
	}
}

func cleanInput(input string) string {
	return strings.ToLower(strings.Trim(input, " "))
}
