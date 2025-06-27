package main

import "fmt"

type Config struct {
	Previous func() error
	Next func()
}

var Page = 0

var urlConfig = Config{
	Previous: func() error {
		if Page <= 1 {
			return fmt.Errorf("you're on the first page")
		}
		Page--
		return nil
	},
	Next: func() {
		Page++
	},
}