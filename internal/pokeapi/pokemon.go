package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonDetails(pokemon string) (RespShallowPokemon ,error) {
	url := baseURL + "/pokemon/" + pokemon
	
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	pokemonDetail := RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokemonDetail)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	
	c.cache.Add(url, dat)
	return pokemonDetail, nil
}