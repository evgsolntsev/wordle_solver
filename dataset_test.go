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
