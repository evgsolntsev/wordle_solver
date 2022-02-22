package main

import "fmt"

type Solver interface {
	GetAttempt() string
	Possibles() []string
	PossiblesCount() int
	ProcessResult(string, []int)
	Reset()
	SetLogLevel(int)
}

type GeneralSolver struct {
	d       *Dataset
	initial *Dataset
	verbose int
}

func NewGeneralSolver(d *Dataset) *GeneralSolver {
	return &GeneralSolver{
		d:       d,
		initial: d.Copy(),
		verbose: 0,
	}
}

func (s *GeneralSolver) PossiblesCount() int {
	return s.d.Len()
}

func (s *GeneralSolver) Possibles() []string {
	return s.d.words
}

func (s *GeneralSolver) ProcessResult(word string, response []int) {
	previousLength := s.PossiblesCount()
	s.d = s.d.Filter(word, response)
	newLength := s.PossiblesCount()
	s.log(fmt.Sprintf("length of dataset: %d => %d", previousLength, newLength), 1)
	if newLength < 100 {
		s.log(fmt.Sprintf("new dataset: %v", s.d.words), 1)
	}
}

func (s *GeneralSolver) Reset() {
	s.d = s.initial.Copy()
}

func (s *GeneralSolver) SetLogLevel(v int) {
	s.verbose = v
}

func (s *GeneralSolver) log(log string, level int) {
	if s.verbose >= level {
		fmt.Println(log)
	}
}

type FrequencyByPlaceSolver struct {
	GeneralSolver
}

var _ Solver = (*FrequencyByPlaceSolver)(nil)

func NewFrequencyByPlaceSolver(d *Dataset) *FrequencyByPlaceSolver {
	general := NewGeneralSolver(d)
	return &FrequencyByPlaceSolver{
		GeneralSolver: *general,
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

type FrequencySolver struct {
	GeneralSolver
}

var _ Solver = (*FrequencySolver)(nil)

func NewFrequencySolver(d *Dataset) *FrequencySolver {
	general := NewGeneralSolver(d)
	return &FrequencySolver{
		GeneralSolver: *general,
	}
}

func (fs *FrequencySolver) CalcFrequency() map[rune]int {
	result := make(map[rune]int)
	for c := 'a'; c <= 'z'; c++ {
		result[c] = 0
	}

	for _, s := range fs.d.words {
		for _, c := range s {
			result[c] += 1
		}
	}

	return result
}

func (fs *FrequencySolver) GetAttempt() string {
	freq := fs.CalcFrequency()
	max := 0
	result := ""
	for _, s := range fs.d.words {
		fs.log(fmt.Sprintf("\nword: %v", s), 2)
		tmp := 0
		found := make(map[rune]bool)
		for i, c := range s {
			if !found[c] {
				tmp += freq[c]
				fs.log(fmt.Sprintf("freq '%v' [%v]: %v", string(c), i, freq[c]), 2)
				found[c] = true
			}
		}
		fs.log(fmt.Sprintf("result: %v", tmp), 2)
		if tmp > max {
			max = tmp
			result = s
		}
	}
	return result
}
