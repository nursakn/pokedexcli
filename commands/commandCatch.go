package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/nursakn/pokedexcli/config"
)

func commandCatch(cfg *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	name := args[0]

	if _, ok := cfg.App.Pokedex[name]; ok {
		return errors.New("you already have this pokemon")
	}
	
	fmt.Println("catching pokemon...")
	res, err := cfg.Client.GetPokemon(strings.ToLower(name))
	
	if err != nil {
		return err
	}

	if rand.Intn(10) < 5 {
		fmt.Println("could not catch(")
		return nil
	}

	fmt.Printf("Caught %s!\n", name)
	cfg.App.Pokedex[name] = res

	return nil
}
