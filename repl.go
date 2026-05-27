package main

import "strings"

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	words := strings.Fields(lower_text)
	return words
}