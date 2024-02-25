package config

import (
	"github.com/nursakn/pokedexcli/pokeapi"
)

type Config struct {
	Client *pokeapi.Api
	App *AppState
}

func InitConfig() *Config {
	api := pokeapi.InitApi(5)
	app := initApp()

	return &Config{
		Client: &api,
		App: &app,
	}
}