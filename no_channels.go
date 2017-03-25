package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	line := "Hello, how is it going? It is nice weather today!"
	cleanedLine := cleanLine(line)
	fmt.Println(makeWordMapFromSlice(cleanedLine))
}

func cleanLine(line string) []string {
	clean := strings.NewReplacer(",", "", ".", "", "?", "", "!", "", "'", "").Replace(line)
	return strings.Split(strings.ToUpper(clean), " ")
}

func makeWordMapFromSlice(line []string) map[string]int {
	wordMap := map[string]int{}
	for _, word := range line {
		wordMap = updateWordMap(wordMap, word)
	}
	return wordMap
}

// END OMIT

func updateWordMap(wordMap map[string]int, word string) map[string]int {
	if _, ok := wordMap[word]; ok {
		wordMap[word]++
	} else {
		wordMap[word] = 1
	}
	return wordMap
}
