package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	counts := make(map[string]int)

	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		cleanWord := strings.Trim(word, ".,!?;:\"'()")
		if cleanWord != "" {
			counts[cleanWord]++
		}
	}

	return counts
}

func main() {
	paragraph := "Go is fast, Go is simple, and learning Go is fun!"

	wordCounts := countWords(paragraph)

	fmt.Println("Word Frequencies:")
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
