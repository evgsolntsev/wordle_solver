package main

import (
	"flag"
)

var (
	datasetFilename string
)

const (
	WORD_LENGTH = 5
	NO_SYMBOL   = 0
	WRONG_PLACE = 1
	CORRECT     = 2
)

func main() {
	flag.Parse()
}

func init() {
	flag.StringVar(&datasetFilename, "dataset", "", "Dataset filename")
}
