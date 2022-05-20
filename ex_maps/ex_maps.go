package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	w_count := make(map[string]int)
	for _, value := range words {
		if _, ok := w_count[value]; ok {
			w_count[value] += 1
		} else {
			w_count[value] = 1
		}
	}
	return w_count
}

func main() {
	wc.Test(WordCount)
}
