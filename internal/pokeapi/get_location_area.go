package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) GetLocationAreas(url string) (*LocationAreaResponse, error) {
	res, err := c.httpClient.Get(url)
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
