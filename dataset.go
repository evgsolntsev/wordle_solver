package main

import (
	"bufio"
	"os"
	"strings"
)

type Dataset struct {
	words []string
}

func NewDataset(words []string) *Dataset {
	newWords := make([]string, 0, len(words))
	for _, word := range words {
		if len(word) == WORD_LENGTH {
			newWords = append(newWords, word)
		}
	}

	return &Dataset{
		words: newWords,
	}
}

func NewDatasetFromFile(filename string) *Dataset {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	words := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		words = append(words, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	return NewDataset(words)
}

func (d *Dataset) Filter(word string, response []int) *Dataset {
	var words []string
	for _, w := range d.words {
		if CheckWord(w, word, response) {
			words = append(words, w)
		}
	}

	return &Dataset{
		words: words,
	}
}

func (d *Dataset) Len() int {
	return len(d.words)
}

func CheckWord(word, attempt string, response []int) bool {
	result := true
	for i := 0; i < WORD_LENGTH; i++ {
		contains := strings.Contains(word, string(attempt[i]))
		equal := word[i] == attempt[i]
		suits := false

		switch response[i] {
		case NO_SYMBOL:
			suits = !contains
		case WRONG_PLACE:
			suits = contains && !equal
		case CORRECT:
			suits = equal
		}

		result = result && suits
	}
	return result
}
