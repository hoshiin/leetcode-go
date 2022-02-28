package main

import (
	"sort"
	"strings"
)

func ReorderLogFiles(logs []string) []string {
	letters := make([]string, 0, len(logs))
	digits := make([]string, 0, len(logs))
	for _, log := range logs {
		words := strings.Split(log, " ")
		if !isDigitsOnly(words[1]) {
			letters = append(letters, log)
			continue
		}
		digits = append(digits, log)
	}
	sort.Slice(letters, func(i, j int) bool {
		ind1 := strings.Index(letters[i], " ")
		ind2 := strings.Index(letters[j], " ")
		if letters[i][ind1+1:] != letters[j][ind2+1:] {
			return letters[i][ind1+1:] < letters[j][ind2+1:]
		}
		return letters[i][:ind1] < letters[j][:ind2]
	})
	return append(letters, digits...)
}

func isDigitsOnly(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
