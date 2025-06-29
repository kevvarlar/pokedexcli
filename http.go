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
    PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
        	URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []any `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func getLocationArea(url string, page int) (string, error) {
	locationAreas := ""
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
		if i == page * 20 {
			locationAreas = fmt.Sprintf("%s%s", locationAreas, locationArea.Name)
		} else {
			locationAreas = fmt.Sprintf("%s%s\n", locationAreas, locationArea.Name)
		}
	}
	return locationAreas, nil
}

func getAllPokemonFromLocationArea(url, name string) (string, error) {
	fullURL := url + name
	pokemons := ""
	req, err := http.Get(fullURL)
	if err != nil {
		return pokemons, fmt.Errorf("failed to get all pokemon: %w", err)
	}
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return pokemons, fmt.Errorf("error while reading request body %w", err)
	}
	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return pokemons, fmt.Errorf("error while unmarshalizing data: %w", err)
	}
	for i, pokemon := range locationArea.PokemonEncounters{
		if i == len(locationArea.PokemonEncounters) - 1 {
			pokemons = fmt.Sprintf("%s - %s", pokemons, pokemon.Pokemon.Name)
		} else {
			pokemons = fmt.Sprintf("%s - %s\n", pokemons, pokemon.Pokemon.Name)
		}
	}
	return pokemons, nil

}

func getPokemonInfo(url, name string) (Pokemon, error) {
	fullURL := url + name
	pokemon := Pokemon{}
	req, err := http.Get(fullURL)
	if err != nil {
		return pokemon, fmt.Errorf("failed to get pokemon info: %w", err)
	}
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return pokemon, fmt.Errorf("error while reading request body: %w", err)
	}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, fmt.Errorf("error while unmarshalizing data: %w", err)
	}
	return pokemon, nil
}