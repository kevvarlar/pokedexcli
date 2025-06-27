package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func(*Config) error
}

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
}

func init() {
    registry["help"] = cliCommand{
        name: "help",
        description: "Displays a help message",
        callback: commandHelp,
    }
}

func commandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, val := range registry {
		fmt.Println(val.name + ": " + val.description)
	}
	return nil
}

func commandMap(c *Config) error {
	c.Next()
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"
	locationAreas, err := getLocationArea(locationAreaURL, Page)
	if err != nil {
		return err
	}
	for _, location := range locationAreas {
		fmt.Println(location)
	}
	return nil
}

func commandMapb(c *Config) error {
	err := c.Previous()
	if err != nil {
		return err
	}
	locationAreaURL := "https://pokeapi.co/api/v2/location-area/"
	locationAreas, err := getLocationArea(locationAreaURL, Page)
	if err != nil {
		return err
	}
	for _, location := range locationAreas {
		fmt.Println(location)
	}
	return nil
}