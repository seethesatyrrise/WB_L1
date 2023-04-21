package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "snow dog sun morning"
	fmt.Println(reversePhrase(phrase))
}

func reversePhrase(phrase string) string {
	words := strings.Fields(phrase)
	if len(words) == 0 {
		return ""
	}
	res := words[0]
	for _, word := range words[1:] {
		res = word + " " + res
	}
	return res
}
