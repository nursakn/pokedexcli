package commands

import (
	"errors"
	"fmt"

	"github.com/nursakn/pokedexcli/config"
)

func commandInspect(cfg *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	name := args[0]

	pokemon, ok := cfg.App.Pokedex[name]

	if !ok {
		return errors.New("you dont have this pokemon yet")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Types:")

	for _, item := range pokemon.Types {
		fmt.Printf("- %v\n", item.Type.Name)
	}

	return nil
}