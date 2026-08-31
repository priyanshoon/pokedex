package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(locationName string) (PokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationName
	if val, ok := c.cache.Get(url); ok {
		pokemonEncounters := PokemonEncounters{}
		err := json.Unmarshal(val, &pokemonEncounters)
		if err != nil {
			return PokemonEncounters{}, err
		}

		return pokemonEncounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncounters{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, err
	}

	pokemonEncounters := PokemonEncounters{}
	err = json.Unmarshal(dat, &pokemonEncounters)
	if err != nil {
		return PokemonEncounters{}, err
	}

	c.cache.Add(url, dat)
	return pokemonEncounters, nil
}

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationsResp := LocationAreas{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
