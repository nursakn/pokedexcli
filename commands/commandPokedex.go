package commands

import (
	"errors"
	"fmt"

	"github.com/nursakn/pokedexcli/config"
)

func commandPokedex(cfg *config.Config, args []string) error {
	if len(cfg.App.Pokedex) == 0 {
		return errors.New("you dont have any pokemons yet")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.App.Pokedex {
		fmt.Printf("- %s \n", pokemon.Name)
	}

	return nil
}