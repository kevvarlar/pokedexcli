package main

import (
	"strings"
	"fmt"
	"os"
	"testing"
	"io"
)

func captureOutput(f func(*Config, string) error, s string) (string, error) {
    orig := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w
    err := f(&urlConfig, s)
    os.Stdout = orig
    w.Close()
    out, _ := io.ReadAll(r)
    return string(out), err
}

func TestCommandHelp(t *testing.T) {
	expectedSubstrings := []string{
		"Welcome to the Pokedex!",
		"Usage:",
		"help: Displays a help message",
		"exit: Exit the Pokedex",
		"map: Displays the next 20 location areas",
		"mapb: Displays the previous 20 location areas",
	}
	out, err := captureOutput(commandHelp, "")
	fmt.Print(out)
	if err != nil {
		t.Errorf("commandHelp returned error: %v\n", err)
	}
	for _, sub := range expectedSubstrings {
		if !strings.Contains(out, sub) {
			t.Errorf("Expected out to include %v\n", sub)
		}
	}
}

func TestCommandMapb_FirstPage(t *testing.T) {
	Page = 1
	err := commandMapb(&urlConfig, "")
	if err == nil {
		t.Error("commandMapb should return error on first page")
	}
	fmt.Printf("Successfully returned error: %v\n", err)
}

func TestCommandMapb_SecondPage(t *testing.T) {
	Page = 2
	expectedPageOne := "canalave-city-area\neterna-city-area\npastoria-city-area\nsunyshore-city-area\nsinnoh-pokemon-league-area\noreburgh-mine-1f\noreburgh-mine-b1f\nvalley-windworks-area\neterna-forest-area\nfuego-ironworks-area\nmt-coronet-1f-route-207\nmt-coronet-2f\nmt-coronet-3f\nmt-coronet-exterior-snowfall\nmt-coronet-exterior-blizzard\nmt-coronet-4f\nmt-coronet-4f-small-room\nmt-coronet-5f\nmt-coronet-6f\nmt-coronet-1f-from-exterior\n"
	out, err := captureOutput(commandMapb, "")
	if err != nil {
		t.Error("commandMapb failed to fetch data for page 1")
	}
	if out != expectedPageOne {
		t.Errorf("Expected %s but got %s", expectedPageOne, out)
	}
}

func TestCommandMap(t *testing.T) {
	Page = 0
	expectedPageOne := "canalave-city-area\neterna-city-area\npastoria-city-area\nsunyshore-city-area\nsinnoh-pokemon-league-area\noreburgh-mine-1f\noreburgh-mine-b1f\nvalley-windworks-area\neterna-forest-area\nfuego-ironworks-area\nmt-coronet-1f-route-207\nmt-coronet-2f\nmt-coronet-3f\nmt-coronet-exterior-snowfall\nmt-coronet-exterior-blizzard\nmt-coronet-4f\nmt-coronet-4f-small-room\nmt-coronet-5f\nmt-coronet-6f\nmt-coronet-1f-from-exterior\n"
	expectedPageTwo := "mt-coronet-1f-route-216\nmt-coronet-1f-route-211\nmt-coronet-b1f\ngreat-marsh-area-1\ngreat-marsh-area-2\ngreat-marsh-area-3\ngreat-marsh-area-4\ngreat-marsh-area-5\ngreat-marsh-area-6\nsolaceon-ruins-2f\nsolaceon-ruins-1f\nsolaceon-ruins-b1f-a\nsolaceon-ruins-b1f-b\nsolaceon-ruins-b1f-c\nsolaceon-ruins-b2f-a\nsolaceon-ruins-b2f-b\nsolaceon-ruins-b2f-c\nsolaceon-ruins-b3f-a\nsolaceon-ruins-b3f-b\nsolaceon-ruins-b3f-c\n"
	out, err := captureOutput(commandMap, "")
	if err != nil {
		t.Error("commandMap failed to fetch data for page 1")
	}
	if out != expectedPageOne {
		t.Errorf("Expected %s but got %s", expectedPageOne, out)
	}
	out, err = captureOutput(commandMap, "")
	if err != nil {
		t.Error("commandMap failed to fetch data for page 2")
	}
	if out != expectedPageTwo {
		t.Errorf("Expected %s but got %s", expectedPageTwo, out)
	}
}

func TestCommandMapCache(t *testing.T) {
	Page = 0
	expectedPageOne := "canalave-city-area\neterna-city-area\npastoria-city-area\nsunyshore-city-area\nsinnoh-pokemon-league-area\noreburgh-mine-1f\noreburgh-mine-b1f\nvalley-windworks-area\neterna-forest-area\nfuego-ironworks-area\nmt-coronet-1f-route-207\nmt-coronet-2f\nmt-coronet-3f\nmt-coronet-exterior-snowfall\nmt-coronet-exterior-blizzard\nmt-coronet-4f\nmt-coronet-4f-small-room\nmt-coronet-5f\nmt-coronet-6f\nmt-coronet-1f-from-exterior\n"
	out, err := captureOutput(commandMap, "")
	if err != nil {
		t.Error("commandMap failed to fetch data for page 1")
	}
	if out != expectedPageOne {
		t.Errorf("Expected %s but got %s", expectedPageOne, out)
	}
	Page = 0
	out, err = captureOutput(commandMap, "")
	if err != nil {
		t.Error("commandMap failed to fetch data for page 1")
	}
	if out != expectedPageOne {
		t.Errorf("Expected %s but got %s", expectedPageOne, out)
	}
}

func TestCommandMapbCache(t *testing.T) {
	Page = 0
	expectedPageOne := "canalave-city-area\neterna-city-area\npastoria-city-area\nsunyshore-city-area\nsinnoh-pokemon-league-area\noreburgh-mine-1f\noreburgh-mine-b1f\nvalley-windworks-area\neterna-forest-area\nfuego-ironworks-area\nmt-coronet-1f-route-207\nmt-coronet-2f\nmt-coronet-3f\nmt-coronet-exterior-snowfall\nmt-coronet-exterior-blizzard\nmt-coronet-4f\nmt-coronet-4f-small-room\nmt-coronet-5f\nmt-coronet-6f\nmt-coronet-1f-from-exterior\n"
	commandMap(&urlConfig, "")
	commandMap(&urlConfig, "")
	out, err := captureOutput(commandMapb, "")
	if err != nil {
		t.Error("commandMap failed to fetch data for page 1")
	}
	if out != expectedPageOne {
		t.Errorf("Expected %s but got %s", expectedPageOne, out)
	}
}

func TestCommandExplore(t *testing.T) {
	cases := []struct {
		locationArea string
		expected []string
	}{
		{
			locationArea: "pastoria-city-area",
			expected: []string{"tentacruel", "wingull", "magikarp"},
		},
		{
			locationArea: "canalave-city-area",
			expected: []string{"staryu", "finneon", "lumineon"},
		},
		{
			locationArea: "mt-coronet-3f",
			expected: []string{"clefairy", "medicham", "solrock"},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			out, err := captureOutput(commandExplore, c.locationArea)
			if err != nil {
				t.Errorf("Error while running explore %v", err)
			}
			for _, s := range c.expected{
				if !strings.Contains(out, s) {
					t.Errorf("Expected output to contain %s", s)
				}
			}
		})
	}
}

func TestCommandExploreFail(t *testing.T) {
	_, err := captureOutput(commandExplore, "")
	if err == nil {
		t.Errorf("Expected an error")
	}

}

func TestCommandExploreCache(t *testing.T) {
	locationArea := "pastoria-city-area"
	expected := []string{"tentacruel", "wingull", "magikarp"}
	out, err := captureOutput(commandExplore, locationArea)
	if err != nil {
		t.Errorf("error while getting pokemon")
	}
	for _, s := range expected {
		if !strings.Contains(out, s) {
			t.Errorf("Expected output to contain %s", s)
		}
	}
	out, err = captureOutput(commandExplore, locationArea)
	if err != nil {
		t.Errorf("error while getting pokemon")
	}
	for _, s := range expected {
		if !strings.Contains(out, s) {
			t.Errorf("Expected output to contain %s", s)
		}
	}
}