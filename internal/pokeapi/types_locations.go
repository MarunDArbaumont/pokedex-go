package pokeapi

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespShallowLocation struct {
	Id int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"location"`
	Name string `json:"name"`
	PokemonEncounters []struct {
		Pokemons struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	
	} `json:"pokemon_encounters"`
}

type RespShallowPokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Height int `json:"height"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct{
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
	BaseExperience int `json:"base_experience"`
}