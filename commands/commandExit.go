package commands

import (
	"os"

	"github.com/nursakn/pokedexcli/config"
)

func commandExit(cfg *config.Config, args []string) error {
	os.Exit(1)
	return nil
}
