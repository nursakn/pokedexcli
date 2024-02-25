package pokeapi

import (
	"net/http"

	"github.com/nursakn/pokedexcli/internal/cache"
)

const (
	locationsURL = "https://pokeapi.co/api/v2/location"
	locationAreaURL = "https://pokeapi.co/api/v2/location-area"
	pokemonsUrl = "https://pokeapi.co/api/v2/pokemon"
)

type Api struct {
	client http.Client
	cache *cache.Cache
}

func InitApi(cacheLife int) Api {
	return Api{
		client: http.Client{},
		cache: cache.NewCache(cacheLife),
	}
}
