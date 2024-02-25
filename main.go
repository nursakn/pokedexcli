package main

import (
	"github.com/nursakn/pokedexcli/config"
)

func main() {
	cfg := config.InitConfig()

	startRepl(cfg)
}
