package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return LocationAreas{}, err
	}

	if res.StatusCode > 299 {
		return LocationAreas{}, errors.New("response failed")
	}

	area := LocationAreas{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		return LocationAreas{}, err
	}

	return area, nil
}
