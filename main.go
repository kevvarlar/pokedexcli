package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		user := cleanInput(scanner.Text())
		command, ok := registry[user[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		var err error
		if len(user) > 1 {
			err = command.callback(&urlConfig, user[1])
		} else {
			err = command.callback(&urlConfig, "")
		}
		if err != nil {
			fmt.Printf("Error while executing command %v\n", err)
		}
	}
}
