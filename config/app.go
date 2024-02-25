package config

import "github.com/nursakn/pokedexcli/pokeapi"

type AppState struct {
	Map MapState
	Pokedex PokedexState
}

type MapState struct {
	PreviousLocationsLink string
	NextLocationsLink string
}

type PokedexState map[string]pokeapi.Pokemon

func initApp() AppState {
	return AppState{
		Map: MapState{
			PreviousLocationsLink: "",
			NextLocationsLink: "",
		},
		Pokedex: make(PokedexState),
	}
}