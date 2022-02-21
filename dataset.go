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
	words := make([]string, len(d.words))
	copy(words, d.words)
	for i := 0; i < WORD_LENGTH; i++ {
		tmpWords := make([]string, 0, len(words))
		for _, w := range words {
			contains := strings.Contains(w, string(word[i]))
			equal := w[i] == word[i]
			suits := false

			switch response[i] {
			case NO_SYMBOL:
				suits = !contains
			case WRONG_PLACE:
				suits = contains && !equal
			case CORRECT:
				suits = equal
			}
			if suits {
				tmpWords = append(tmpWords, w)
			}
		}
		words = tmpWords
	}

	return &Dataset{
		words: words,
	}
}

func (d *Dataset) Len() int {
	return len(d.words)
}
