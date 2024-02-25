package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (api *Api) MapLocations(link string) (RespShallowLocations, error) {
	url := locationsURL

	if link != "" {
		url = link
	}

	locations := RespShallowLocations{} 

	val, ok := api.cache.Get(url)

	if ok {
		err := json.Unmarshal(val, &locations)
		return locations, err
	}

	resp, err := api.client.Get(url)

	if err != nil {
		return locations, err
	}

  data, err := io.ReadAll(resp.Body)

	if err != nil {
		return locations, err
	}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return locations, err
	}

	api.cache.Add(url, data)

	return locations, err
}

func (api *Api) GetLocation(name string) (Location, error) {
	location := Location{}
	
	url := fmt.Sprintf("%s/%s", locationAreaURL, name)

	val, ok := api.cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &location)
		return location, err
	}

	resp, err := api.client.Get(url)

	if err != nil {
		return location, err
	}

  data, err := io.ReadAll(resp.Body)

	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)

	if err != nil {
		return location, err
	}

	api.cache.Add(url, data)
	return location, nil
}
