package main

import (
	"fmt"
	"strings"
)

func main() {
	var genetic_code = "AATCCGCTAGATCGATCTCCAGGATTTATACCGGTACA"
	if !basesOnly(genetic_code) {
		fmt.Println("Please input a valid genetic sequence.")
	} else {
		sequence(genetic_code)
	}
}

// dna uses 't' while rna uses 'u'
const bases = "acgt"

func basesOnly(s string) bool {
	for _, char := range s {
		if !strings.Contains(bases, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func sequence(s string) {
	type sequence struct {
		codon string
		count int
	}

	var m = make(map[string]int)

	sequence_len := len(s)
	sequence_codons := len(s) - 2
	for sequence_codons > 0 {
		m[s[sequence_codons-1:sequence_len]] += 1
		sequence_len -= 1
		sequence_codons -= 1
	}
	fmt.Println(m)
}
