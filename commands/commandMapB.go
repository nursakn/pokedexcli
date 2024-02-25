package commands

import (
	"errors"
	"fmt"

	"github.com/nursakn/pokedexcli/config"
)
	

func commandMapb(cfg *config.Config, args []string) error {
	if cfg.App.Map.PreviousLocationsLink == "" {
		return errors.New("you are on first page")
	}

	res, err := cfg.Client.MapLocations(cfg.App.Map.PreviousLocationsLink)

	if err != nil {
		return err
	}

	if res.Next != nil {
		cfg.App.Map.NextLocationsLink = *res.Next
	} else {
		cfg.App.Map.NextLocationsLink = ""
	}
	
	if res.Previous != nil {
		cfg.App.Map.PreviousLocationsLink = *res.Previous
	} else {
		cfg.App.Map.PreviousLocationsLink = ""
	}

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}

	return err
}
