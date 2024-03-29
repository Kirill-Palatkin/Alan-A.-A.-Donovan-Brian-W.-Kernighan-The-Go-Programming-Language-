package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "ch03/ex12: must have 2 arguments.")
		os.Exit(1)
	}
	fmt.Printf("%t\n", anagram(os.Args[1], os.Args[2]))
}

func anagram(a, b string) bool {
	return equals(runeOccurrences(a), runeOccurrences(b))
}

func runeOccurrences(s string) map[rune]int {
	occurrences := make(map[rune]int)
	for _, r := range s {
		occurrences[r]++
	}
	return occurrences
}

func equals(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}


//go run main.go anagrams arsmagna
