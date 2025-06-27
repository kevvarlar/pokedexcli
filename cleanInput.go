package main

import "strings"

func cleanInput(text string) []string {
	result := []string{}
	text = strings.ToLower(strings.Trim(text, " "))
	currentString := ""
	if len(text) == 0 {
		return []string{""}
	}
	for i, c := range text {
		if i == len(text) - 1 {
			result = append(result, currentString + string(c))
		}
		if string(c) == " " {
			if currentString == "" {
				continue
			}
			result = append(result, currentString)
			currentString = ""
			continue
		}
		currentString += string(c)
	}
	return result
}