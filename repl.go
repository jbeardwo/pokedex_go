package main

import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	out := strings.Fields(text)
	return out
}
