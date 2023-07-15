package main

import (
	"strings"
	"unicode"
)

// the tokenizer is the first step of text normalisation.
// its job is to convert text into a list of tokens.
// our implementation splits the text on a word
// boundary and removes punctuation marks:
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func normalize(text string) []string {
	tokens := tokenize(text)
	tokens = languageNormalizer(tokens)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
