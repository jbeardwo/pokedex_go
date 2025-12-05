package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationAreas(url string) (*LocationAreaResponse, error) {
	data, ok := c.pokeCache.Get(url)
	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return &LocationAreaResponse{}, err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return &LocationAreaResponse{}, fmt.Errorf("Status Code: %v", res.StatusCode)
		}
		if err != nil {
			return &LocationAreaResponse{}, err
		}
		c.pokeCache.Add(url, data)
	}
	locationRes := &LocationAreaResponse{}
	err := json.Unmarshal(data, &locationRes)
	if err != nil {
		return &LocationAreaResponse{}, err
	}

	return locationRes, nil
}
