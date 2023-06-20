package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	//ls := strings.Fields(s)
	//For range and gettiing the value of words
	for _, w := range strings.Fields(s) {
		_, isPresent := m[w]
		if isPresent {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}
	return m
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
