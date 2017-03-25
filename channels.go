package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	line := "Hello, how is it going? It is nice weather today!"
	wordChannel := getCleanedWords(line)
	fmt.Println(makeWordMapFromChan(wordChannel))
}

func getCleanedWords(line string) chan string {
	wordChan := make(chan string)
	go func() { // HL
		for _, word := range cleanWords(line) { // HL
			wordChan <- word // HL
		} // HL
		close(wordChan) // HL
	}() // HL
	return wordChan
}

func makeWordMapFromChan(wordChan chan string) map[string]int {
	wordMap := map[string]int{}
	for word := range wordChan { // HL
		wordMap = updateWordMap(wordMap, word) // HL
	} // HL
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

func cleanWords(line string) []string {
	return strings.Split(strings.ToUpper(strings.NewReplacer(",", "", ".", "", "?", "", "!", "", "'", "").Replace(line)), " ")
}
