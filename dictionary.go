package main

import "unicode"
import "math/rand"
import "bufio"

type Dictionary struct {
	index []rune
	words map[rune][]string
}

func MakeDictionary(scanner *bufio.Scanner) (*Dictionary, error) {
	dict := Dictionary {words:make(map[rune][]string)}

	// Add words
	for scanner.Scan() {
		word := scanner.Text()
		first := []rune(word)[0]
		dict.words[first] = append(dict.words[first], word)
	}
	
	// Build index
	for f := range dict.words {
		dict.index = append(dict.index, f)
	}

	return &dict, scanner.Err()
}

func (dict *Dictionary) RandomWord(first rune) string {
    firstLower := unicode.ToLower(first)
    answers := dict.words[firstLower]
    return answers[rand.Intn(len(answers))]
}

func (dict *Dictionary) RandomFirst() rune {
    return dict.index[rand.Intn(len(dict.index))]
}
