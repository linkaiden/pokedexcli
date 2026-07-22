package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationInfo(area_name string) (Location, error) {
	url := baseURL + "/location-area/" + area_name

	if val, ok := c.cache.Get(url); ok {
		locationInfo := Location{}
		err := json.Unmarshal(val, &locationInfo)
		if err != nil {
			return Location{}, err
		}
		return locationInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationInfo := Location{}
	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		return Location{}, err
	}

	// Add to cache after confirmed good API data can be unmarshalled.
	c.cache.Add(url, data)

	return locationInfo, nil
}
