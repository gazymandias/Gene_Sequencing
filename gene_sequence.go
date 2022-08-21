package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
	
)

func main() {
	start := time.Now()
	n := 0
	for n < 1000000 {
		// var genetic_code = "AATCCGCTAGATCGATCTCCAGGATTTATACCGGTACA"
		var genetic_code = DummyGeneticCode(1000)
		if !basesOnly(genetic_code) {
			fmt.Println("Please input a valid genetic sequence.")
		} else {
			go sequence(genetic_code)
		}
		n += 1
	}
	elapsed := time.Since(start)
	log.Printf("Process took %s", elapsed)
}

// dna uses 't' while rna uses 'u'
const bases = "acgt"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func DummyGeneticCode(length int) string {
	return StringWithCharset(length, bases)
}

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
