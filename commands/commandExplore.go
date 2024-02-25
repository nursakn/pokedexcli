package commands

import (
	"errors"
	"fmt"

	"github.com/nursakn/pokedexcli/config"
)

func commandExplore(cfg *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	name := args[0]

	fmt.Println("exploring location...")
	res, err := cfg.Client.GetLocation(name)
	
	if err != nil {
		return err
	}

	fmt.Println("Encountered pokemons:")
	for _, item := range res.PokemonEncounters {
		fmt.Printf("- %s\n", item.Pokemon.Name)
	}

	return nil
}