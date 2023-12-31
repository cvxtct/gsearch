package main

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

// to make the search case-insensitive
func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// dropping common words
// almost any English text contains commonly used words like
// a, I, the or be. Such words are called stop words.
// e are going to remove them since almost any document would match the stop words
// here is no "official" list of stop words
// let's exclude the top 10 by the OEC rank. Feel free to add more
func (p *Project)stopwordFilter(tokens []string) []string {
	// var stopwords = map[string]struct{}{
	// 	// TODO extend -> so so
	// 	// TODO .md specific simbols??
	// 	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	// 	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	// }

	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := p.stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

// Stemming
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}

// TODO
// choose for removal set of special characters -> collect them
// ignore non english words and non english characters
// review and extend stop words
// recognise urls, email adresses (?)
