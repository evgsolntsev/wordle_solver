package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcFrequency(t *testing.T) {
	for i, testcase := range []struct {
		words   []string
		results map[rune][]int
	}{{
		words: []string{"abaca", "babab", "cbacb"},
		results: map[rune][]int{
			'a': []int{1, 1, 2, 1, 1},
			'b': []int{1, 2, 1, 0, 2},
			'c': []int{1, 0, 0, 2, 0},
		},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			freq := NewFrequencyByPlaceSolver(NewDataset(testcase.words)).CalcFrequency()
			for k, v := range testcase.results {
				assert.Equal(t, v, freq[k])
			}
			for k, v := range freq {
				if _, ok := testcase.results[k]; !ok {
					assert.Equal(t, []int{0, 0, 0, 0, 0}, v)
				}
			}
		})
	}
}

func TestGetAttempt(t *testing.T) {
	for i, testcase := range []struct {
		words    []string
		result   string
	}{{
		words:    []string{"abaac", "babac", "abaca", "xandr", "abbac", "ataac"},
		result:   "abaac",
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			solver := NewFrequencyByPlaceSolver(NewDataset(testcase.words))
			//solver.verbose = 2
			assert.Equal(t, testcase.result, solver.GetAttempt())
		})
	}
}
