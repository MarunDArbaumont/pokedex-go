package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client)ListLocations(pageURL) (RespShallowLocations ,error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	} else {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	location := RespShallowLocations{}
	err = json.Unmarshal(dat, &location)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return location, nil
}