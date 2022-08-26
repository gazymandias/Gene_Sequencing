package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	n := 0
	for n < 10000000 {
		// var genetic_code = "cgcagacgcgcaagcagattttataaagccaatgtatatacaactcctaccaggccaaacaaggtccgcgcacacgaaacctggaaaactttgtgtcggc"
		var genetic_code = DummyGeneticCode(50)
		if !basesOnly(genetic_code) {
			fmt.Println("Please input a valid genetic sequence.")
		} else {
			go fmt.Println(sequence(genetic_code))
		}
		n += 1
	}
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

func findSequenceRange(s string) (int, int, error) {
	start_codon := strings.Index(s, "atg")
	stop_codon_one := strings.Index(s, "taa")
	stop_codon_two := strings.Index(s, "tga")
	stop_codon_three := strings.Index(s, "taa")
	stop_codons := []int{stop_codon_one, stop_codon_two, stop_codon_three}
	stop_codon := stop_codons[0]
	for _, v := range stop_codons {
		if (v < stop_codon || stop_codon < 0) && v > start_codon {
			stop_codon = v
		}
	}
	if start_codon == -1 {
		return -1, -1, errors.New("no valid start codon")
	}
	if stop_codon < start_codon {
		return -1, -1, errors.New("start codon is after stop codon")
	}
	if stop_codon == 0 {
		return -1, -1, errors.New("no valid stop codon after start codon")
	}
	return start_codon, stop_codon, nil
}

func cleanSequenceRange(s string) (string, error) {
	start_codon, stop_codon, err := findSequenceRange(s)
	if err != nil {
		return s, err
	}
	s = s[start_codon:][0 : (stop_codon+3)-start_codon]
	return s, nil
}

func sequence(s string) (string, map[string]int, error) {
	type sequence struct {
		codon string
		count int
	}
	var o = s
	var m = make(map[string]int)
	s, err := cleanSequenceRange(s)
	if err != nil {
		return o, m, err
	}
	var codons_to_sequence int = 0
	if len(s)%3 == 0 {
		codons_to_sequence = len(s) / 3
	} else {
		codons_to_sequence = (len(s) - (len(s) % 3)) / 3
		s = s[:len(s)-len(s)%3]
	}
	for codons_to_sequence > 0 {
		m[s[:3]] += 1
		s = s[3:]
		codons_to_sequence -= 1
	}
	return o, m, err
}
