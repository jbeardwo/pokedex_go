package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(url string) (*Pokemon, error) {
	data, ok := c.pokeCache.Get(url)
	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return &Pokemon{}, err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return &Pokemon{}, fmt.Errorf("Status Code: %v", res.StatusCode)
		}
		if err != nil {
			return &Pokemon{}, err
		}
		c.pokeCache.Add(url, data)
	}
	pokemonRes := &Pokemon{}
	err := json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return &Pokemon{}, err
	}

	return pokemonRes, nil
}
