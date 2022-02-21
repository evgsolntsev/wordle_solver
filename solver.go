package main

import "fmt"

type FrequencyByPlaceSolver struct {
	d *Dataset
	verbose bool
}

func NewFrequencyByPlaceSolver(d *Dataset) *FrequencyByPlaceSolver {
	return &FrequencyByPlaceSolver{
		d: d,
		verbose: false,
	}
}

func (s *FrequencyByPlaceSolver) CalcFrequency() map[rune][]int {
	result := make(map[rune][]int)
	for c := 'a'; c <= 'z'; c++ {
		result[c] = make([]int, WORD_LENGTH)
	}

	for _, s := range s.d.words {
		for i, c := range s {
			result[c][i] += 1
		}
	}

	return result
}

func (fs *FrequencyByPlaceSolver) GetAttempt() string {
	freq := fs.CalcFrequency()
	max := 0
	result := ""
	for _, s := range fs.d.words {
		
		fs.log(fmt.Sprintf("\nword: %v", s))
		tmp := 0
		for i, c := range s {
			tmp += freq[c][i]
			fs.log(fmt.Sprintf("freq '%v' [%v]: %v", string(c), i, freq[c][i]))
		}
		fs.log(fmt.Sprintf("result: %v", tmp))
		if tmp > max {
			max = tmp
			result = s
		}
	}
	return result
}

func (s *FrequencyByPlaceSolver) log(log string) {
	if s.verbose {
		fmt.Println(log)
	}
}
