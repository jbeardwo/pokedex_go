package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetExploreArea(url string) (*ExploreAreaResponse, error) {
	data, ok := c.pokeCache.Get(url)
	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return &ExploreAreaResponse{}, err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return &ExploreAreaResponse{}, fmt.Errorf("Status Code: %v", res.StatusCode)
		}
		if err != nil {
			return &ExploreAreaResponse{}, err
		}
		c.pokeCache.Add(url, data)
	}
	areaRes := &ExploreAreaResponse{}
	err := json.Unmarshal(data, &areaRes)
	if err != nil {
		return &ExploreAreaResponse{}, err
	}

	return areaRes, nil
}
