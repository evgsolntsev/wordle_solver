package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	datasetFilename    string
	mode               string
	solverType             string
	verbose            bool
	POSSIBLE_RESPONSES = []int{NO_SYMBOL, WRONG_PLACE, CORRECT}
)

const (
	WORD_LENGTH               = 5
	NO_SYMBOL                 = 0
	WRONG_PLACE               = 1
	CORRECT                   = 2
	CHECKER_MODE              = "c"
	INTERACTIVE_MODE          = "i"
	FREQUENCY_BY_PLACE_SOLVER = "freqByPlace"
	FREQUENCY_SOLVER          = "freq"
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

	var solver Solver
	switch solverType {
	case FREQUENCY_BY_PLACE_SOLVER:
		solver = NewFrequencyByPlaceSolver(dataset)
	case FREQUENCY_SOLVER:
		solver = NewFrequencySolver(dataset)
	}

	if verbose {
		solver.SetLogLevel(1)
	}

	switch mode {
	case INTERACTIVE_MODE:
		interactive(solver)
	case CHECKER_MODE:
		check(dataset, solver)
	default:
		fmt.Printf("Unknown mode '%s'.\n", mode)
	}
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

func check(dataset *Dataset, solver Solver) {
	CHECK_WORDS_LEN := 2000
	hiddenWords := make([]string, 0, CHECK_WORDS_LEN)
	for i := 0; i < CHECK_WORDS_LEN; i++ {
		hiddenWords = append(hiddenWords, dataset.words[rand.Intn(len(dataset.words))])
	}

	attemptsSummary := 0
	for _, hiddenWord := range hiddenWords {
		solver.Reset()
		result := ""
		attempts := make([]string, 0)

		for {
			if solver.PossiblesCount() == 0 {
				result = "FAIL"
				break
			}

			attempt := solver.GetAttempt()
			attempts = append(attempts, attempt)
			response := GenerateResponse(hiddenWord, attempt)
			solver.ProcessResult(attempt, response)

			success := true
			for _, r := range response {
				if r != CORRECT {
					success = false
				}
			}

			if success {
				result = fmt.Sprintf("OK (%d attempts)", len(attempts))
				break
			}
		}
		fmt.Printf("%s: %s %v\n", hiddenWord, result, attempts)

		attemptsSummary += len(attempts)
	}
	fmt.Println()
	fmt.Printf("Average attempts number: %v\n", float64(attemptsSummary)/float64(CHECK_WORDS_LEN))
}

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.StringVar(&datasetFilename, "dataset", "dataset.txt", "Dataset filename")
	flag.StringVar(&mode, "mode", "i", "Work mode")
	flag.StringVar(&solverType, "solver", FREQUENCY_BY_PLACE_SOLVER, "Solver")
	flag.BoolVar(&verbose, "v", false, "Verbose logging")
}
