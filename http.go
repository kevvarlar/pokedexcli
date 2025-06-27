package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
    ID                   int    `json:"id"`
    Name                 string `json:"name"`
    GameIndex            int    `json:"game_index"`
    EncounterMethodRates []any  `json:"encounter_method_rates"`
    Location             struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"location"`
    Names             []any `json:"names"`
    PokemonEncounters []any `json:"pokemon_encounters"`
}
func getLocationArea(url string, page int) ([]string, error) {
	locationAreas := []string{}
	for i := page * 20 - 19; i < page * 20 + 1; i++ {
		fullURL := fmt.Sprintf("%v%v", url, i)
		req, err := http.Get(fullURL)
		if err != nil {
			return locationAreas, fmt.Errorf("failed to get all location areas: %w", err)
		}
		defer req.Body.Close()
		data, err := io.ReadAll(req.Body)
		if err != nil {
			return locationAreas, fmt.Errorf("error while reading request body %w", err)
		}
		var locationArea LocationArea
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return locationAreas, fmt.Errorf("error while unmarshalizing data: %w", err)
		}
		locationAreas = append(locationAreas, locationArea.Name)
	}
	return locationAreas, nil
}