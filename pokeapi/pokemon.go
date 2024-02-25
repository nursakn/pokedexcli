package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (api *Api) GetPokemon(pokemonName string) (Pokemon, error) {
	pokemon := Pokemon{}

	url := fmt.Sprintf("%s/%s", pokemonsUrl, pokemonName)

	val, ok := api.cache.Get(url)

	if ok {
		err := json.Unmarshal(val, &pokemon)
		return pokemon, err
	}

	resp, err := api.client.Get(url)

	if err != nil {
		return pokemon, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return pokemon, err
	}

	api.cache.Add(url, data)

	return pokemon, nil
}
