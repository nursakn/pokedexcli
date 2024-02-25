package commands

import (
	"fmt"

	"github.com/nursakn/pokedexcli/config"
)

func commandHelp(cfg *config.Config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range GetCommands() {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}