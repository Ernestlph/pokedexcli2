package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonname string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonname
	fullUrl := baseURL + endpoint

	// Check if the data is in the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit!")
		pokemonDetail := Pokemon{}
		err := json.Unmarshal(data, &pokemonDetail)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonDetail, nil
	}

	// If Pokemon not found in cache, make a request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonDetail := Pokemon{}
	err = json.Unmarshal(dat, &pokemonDetail)

	if err != nil {
		return Pokemon{}, err
	}

	// Add the data to the cache
	c.cache.Add(fullUrl, dat)

	return pokemonDetail, nil

}
