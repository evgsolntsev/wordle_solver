package main

import (
	"flag"
	"fmt"
)

var (
	datasetFilename    string
	verbose            bool
	POSSIBLE_RESPONSES = []int{NO_SYMBOL, WRONG_PLACE, CORRECT}
)

const (
	WORD_LENGTH = 5
	NO_SYMBOL   = 0
	WRONG_PLACE = 1
	CORRECT     = 2
)

func main() {
	flag.Parse()

	if verbose {
		fmt.Printf("downloading dataset from '%s'...\n", datasetFilename)
	}
	dataset := NewDatasetFromFile(datasetFilename)
	if verbose {
		fmt.Printf("dataset length: %d\n", len(dataset.words))
		fmt.Println("first word:", dataset.words[0])
		fmt.Println("last word:", dataset.words[len(dataset.words)-1])
		fmt.Println("loaded.")
		fmt.Println()
	}

	solver := NewFrequencyByPlaceSolver(dataset)
	if verbose {
		solver.verbose = 1
	}
	interactive(solver)
}

func interactive(solver Solver) {
	fmt.Println("Type 'exit' if you want to stop.")
	var input string
	for {
		attempt := solver.GetAttempt()
		fmt.Printf("Is it '%s'?\n> ", attempt)

		var response []int
		for {
			fmt.Scanln(&input)
			if input == "exit" {
				fmt.Println("Well, bye.")
				return
			}
			if len(input) != WORD_LENGTH {
				fmt.Printf("I don't understand. Repeat please.\n> ")
				continue
			}

			response = make([]int, 0, WORD_LENGTH)
			for _, c := range input {
				for _, possible := range POSSIBLE_RESPONSES {
					if c == rune(fmt.Sprintf("%v", possible)[0]) {
						response = append(response, possible)
					}
				}
			}
			if len(response) == WORD_LENGTH {
				break
			}
			fmt.Printf("I don't understand. Repeat please.\n.> ")
		}

		done := true
		for _, r := range response {
			if r != CORRECT {
				done = false
			}
		}
		if done {
			fmt.Println("Hooray!")
			return
		}

		solver.ProcessResult(attempt, response)
		if solver.PossiblesCount() == 0 {
			fmt.Println("Sorry, I don't know such a word.")
			return
		}
	}
}

func init() {
	flag.StringVar(&datasetFilename, "dataset", "dataset.txt", "Dataset filename")
	flag.BoolVar(&verbose, "v", false, "Verbose logging")
}
