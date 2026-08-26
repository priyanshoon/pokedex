package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

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
