package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (*LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return &LocationAreaResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return &LocationAreaResponse{}, fmt.Errorf("Status Code: %v", res.StatusCode)
	}
	if err != nil {
		return &LocationAreaResponse{}, err
	}

	locationRes := &LocationAreaResponse{}
	err = json.Unmarshal(body, &locationRes)
	return locationRes, nil
}
