package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand"
	"github.com/kevvarlar/pokedexcli/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(*Config, string) error
}

type Pokemon struct {
    ID int `json:"id"`
    Name string `json:"name"`
    BaseExperience int `json:"base_experience"`
    Height int `json:"height"`
    IsDefault bool `json:"is_default"`
    Order int `json:"order"`
    Weight int `json:"weight"`
    Abilities []any `json:"abilities"`
    Forms []any `json:"forms"`
    GameIndices []any `json:"game_indices"`
    HeldItems []any `json:"held_items"`
    LocationAreaEncounters string `json:"location_area_encounters"`
    Moves []any `json:"moves"`
    PastTypes []any `json:"past_types"`
    PastAbilities []any `json:"past_abilities"`
    Sprites any `json:"sprites"`
    Cries any `json:"cries"`
    Species any `json:"species"`
    Stats []struct{
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct{
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
    Types []struct{
		Slot int `json:"slot"`
		Type struct{
			Name string `json:"name"`
			URL string `json:"url"`
		}
	} `json:"types"`
}

var cache = pokecache.NewCache(5 * time.Second)

var pokedex = make(map[string]Pokemon)

var registry = map[string]cliCommand{
	"help" : {
		name: "help",
		description: "Displays a help message",
		callback: nil,
	},
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"map": {
		name: "map",
		description: "Displays the next 20 location areas",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays the previous 20 location areas",
		callback: commandMapb,
	},
	"explore": {
		name: "explore",
		description: "Displays all pokemon from given location area, i.e. explore [location_area]",
		callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Catch a pokemon, chance of catching is based on base experience, i.e. catch [pokemon_name]",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "Inspect a pokemon's info, i.e inspect [pokemon_name]",
		callback: commandInspect,
	},
}

func init() {
    registry["help"] = cliCommand{
        name: "help",
        description: "Displays a help message",
        callback: commandHelp,
    }
}

func commandExit(*Config, string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config, string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, val := range registry {
		fmt.Println(val.name + ": " + val.description)
	}
	return nil
}

func commandMap(c *Config, _ string) error {
	c.Next()
	v, ok := cache.Get(strconv.Itoa(Page))
	if ok {
		fmt.Println(string(v))
		return nil
	}
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"
	locationAreas, err := getLocationArea(locationAreaURL, Page)
	if err != nil {
		return err
	}
	cache.Add(strconv.Itoa(Page), []byte(locationAreas))
	fmt.Println(locationAreas)
	return nil
}

func commandMapb(c *Config, _ string) error {
	err := c.Previous()
	if err != nil {
		return err
	}
	v, ok := cache.Get(strconv.Itoa(Page))
	if ok {
		fmt.Println(string(v))
		return nil
	}
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"
	locationAreas, err := getLocationArea(locationAreaURL, Page)
	if err != nil {
		return err
	}
	fmt.Println(locationAreas)
	return nil
}

func commandExplore(_ *Config, locationArea string) error {
	if len(locationArea) == 0 {
		return fmt.Errorf("no location area given")
	}
	v, ok := cache.Get(locationArea)
	if ok {
		fmt.Printf("Exploring %s...\n", locationArea)
		fmt.Println("Found Pokemon:")
		fmt.Println(string(v))
		return nil
	}
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"
	pokemons, err := getAllPokemonFromLocationArea(locationAreaURL, locationArea)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationArea)
	fmt.Println("Found Pokemon:")
	fmt.Println(pokemons)
	cache.Add(locationArea, []byte(pokemons))
	return nil
}

func catchPokemon(pokemonName string, pokemon Pokemon) {
	fmt.Println("Throwing a Pokeball at", pokemonName + "...")
	chance := rand.Float32() * 10
	base := float32(pokemon.BaseExperience) / 35.0
	if chance - base > 0 {
		fmt.Println(pokemonName, "was caught!")
		pokedex[pokemonName] = pokemon
	} else {
		fmt.Println(pokemonName, "escaped!")
	}
}

func commandCatch(_ *Config, pokemonName string) error {

	pokemonURL := "https://pokeapi.co/api/v2/pokemon/"
	v, ok := cache.Get(pokemonName)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(v, &pokemon)
		if err != nil {
			return err
		}
		catchPokemon(pokemonName, pokemon)
		return nil
	}
	pokemon, err := getPokemonInfo(pokemonURL, pokemonName)
	if err != nil {
		return err
	}
	catchPokemon(pokemonName, pokemon)

	//Cache logic
	pokemonJSON, err := json.Marshal(pokemon)
	if err != nil {
		return err
	}
	cache.Add(pokemonName, pokemonJSON)
	return nil
}

func commandInspect(_ *Config, pokemonName string) error {
	pokemon, ok := pokedex[pokemonName]
	fmt.Println("here")
	if !ok {
		return fmt.Errorf("Pokemon not in Pokedex")
	}
	fmt.Println("Name:", pokemon.Name, "\nHeight:", pokemon.Height, "\nWeight:", pokemon.Weight, "\nStats:")
	for _, stat := range pokemon.Stats{
		fmt.Println("  -" + stat.Stat.Name + ":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types{
		fmt.Println("  -", t.Type.Name)
	}
	return nil
}