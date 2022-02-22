package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	for i, testcase := range []struct {
		words    []string
		word     string
		response []int
		result   []string
	}{{
		words:    []string{"abaac", "babac", "abaca", "xandr", "abbac", "ataac"},
		word:     "xaabd",
		response: []int{0, 1, 2, 1, 0}, // hidden word is "abaca"
		result:   []string{"abaac", "abaca"},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := NewDataset(testcase.words).Filter(testcase.word, testcase.response)
			assert.Len(t, result.words, len(testcase.result))
			for _, w := range testcase.result {
				assert.Contains(t, testcase.result, w)
			}
		})
	}
}

func TestCheckWord(t *testing.T) {
	for i, testcase := range []struct {
		word     string
		attempt  string
		response []int
		result   bool
	}{{
		word:     "xaabd",
		attempt:  "yyyyy",
		response: []int{0, 0, 0, 0, 0},
		result:   true,
	}, {
		word:     "abaca",
		attempt:  "baaac",
		response: []int{1, 1, 2, 1, 1},
		result:   true,
	}, {
		word:     "offer",
		attempt:  "order",
		response: []int{2, 0, 0, 2, 2},
		result:   true,
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.Equal(t, testcase.result, CheckWord(
				testcase.word, testcase.attempt, testcase.response), testcase.result)
		})
	}
}

func TestGenerateResponse(t *testing.T) {
	for i, testcase := range []struct {
		word     string
		attempt  string
		response []int
	}{{
		word:     "halve",
		attempt:  "sores",
		response: []int{0, 0, 0, 1, 0},
	}, {
		word:     "rooms",
		attempt:  "spite",
		response: []int{1, 0, 0, 0, 0},
	}} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.Equal(t, testcase.response, GenerateResponse(testcase.word, testcase.attempt))
		})
	}

}
