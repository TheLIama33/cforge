package tokenizer

import (
	"strings"
	"unicode/utf8"
)

type Estimator interface {
	Count(text string) int
}

type SimpleEstimator struct{}

func (e SimpleEstimator) Count(text string) int {
	if text == "" {
		return 0
	}

	chars := utf8.RuneCountInString(text)
	return chars / 4
}

func Estimate(text string) int {
	length := len(text)
	if length == 0 {
		return 0
	}

	return length / 4
}

func CountLines(text string) int {
	return strings.Count(text, "\n") + 1
}
