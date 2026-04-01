package pokeapi

import (
	"net/http"
	"time"

	"github.com/MarunDArbaumont/pokedex-go/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	Pokedex map[string]RespShallowPokemon
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		Pokedex: make(map[string]RespShallowPokemon),
	}
}