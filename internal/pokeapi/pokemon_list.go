package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDetails(location string) (RespShallowLocation ,error) {
	url := baseURL + "/location-area/" + location
	
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}
	pokemons := RespShallowLocation{}
	err = json.Unmarshal(dat, &pokemons)
	if err != nil {
		return RespShallowLocation{}, err
	}
	
	c.cache.Add(url, dat)
	return pokemons, nil
}