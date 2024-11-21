package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}

	// Check if the data is in the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	// Add the data to the cache
	c.cache.Add(fullUrl, dat)

	return locationAreasResp, nil

}

func (c *Client) ListLocationAreaDetail(locationAreaName string) (LocationAreaRespDetail, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseURL + endpoint

	// Check if the data is in the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		fmt.Println("Cache hit!")
		locationAreaDetail := LocationAreaRespDetail{}
		err := json.Unmarshal(data, &locationAreaDetail)
		if err != nil {
			return LocationAreaRespDetail{}, err
		}
		return locationAreaDetail, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaRespDetail{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRespDetail{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationAreaRespDetail{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaRespDetail{}, err
	}

	locationAreaDetail := LocationAreaRespDetail{}
	err = json.Unmarshal(dat, &locationAreaDetail)

	if err != nil {
		return LocationAreaRespDetail{}, err
	}

	// Add the data to the cache
	c.cache.Add(fullUrl, dat)

	return locationAreaDetail, nil

}
