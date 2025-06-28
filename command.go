package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kevvarlar/pokedexcli/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(*Config, string) error
}

var cache = pokecache.NewCache(5 * time.Second)

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