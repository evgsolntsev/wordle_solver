package main

import "fmt"

type Solver interface {
	GetAttempt() string
	PossiblesCount() int
	ProcessResult(string, []int)
}

type FrequencyByPlaceSolver struct {
	d *Dataset
	verbose int
}

var _ Solver = (*FrequencyByPlaceSolver)(nil)

func NewFrequencyByPlaceSolver(d *Dataset) *FrequencyByPlaceSolver {
	return &FrequencyByPlaceSolver{
		d: d,
		verbose: 0,
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
		
		fs.log(fmt.Sprintf("\nword: %v", s), 2)
		tmp := 0
		for i, c := range s {
			tmp += freq[c][i]
			fs.log(fmt.Sprintf("freq '%v' [%v]: %v", string(c), i, freq[c][i]), 2)
		}
		fs.log(fmt.Sprintf("result: %v", tmp), 2)
		if tmp > max {
			max = tmp
			result = s
		}
	}
	return result
}

func (s *FrequencyByPlaceSolver) PossiblesCount() int {
	return s.d.Len()
}

func (s *FrequencyByPlaceSolver) ProcessResult(word string, response []int) {
	previousLength := s.PossiblesCount()
	s.d = s.d.Filter(word, response)
	newLength := s.PossiblesCount()
	s.log(fmt.Sprintf("length of dataset: %d => %d", previousLength, newLength), 1)
}

func (s *FrequencyByPlaceSolver) log(log string, level int) {
	if s.verbose >= level {
		fmt.Println(log)
	}
}
