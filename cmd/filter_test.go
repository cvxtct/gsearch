package main

import (
	"fmt"
	"reflect"
	"testing"
)

var proj Project

func TestLowercaseFilter(t *testing.T) {
	var lowercaseFilterTest = []struct {
		tokens   []string
		expected []string
	}{
		{tokens: []string{"FOO", "Bar", "baZ", "oNe"}, expected: []string{"foo", "bar", "baz", "one"}},
	}

	// TODO - FIXME
	proj.stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	for _, tt := range lowercaseFilterTest {
		testname := fmt.Sprintf("%v", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			actual := lowercaseFilter(tt.tokens)
			loCased := reflect.DeepEqual(actual, tt.expected)
			if !loCased {
				t.Errorf("got %v expected %v", actual, tt.expected)
			}
		})
	}
}

func TestStopwordFilter(t *testing.T) {
	var stopwordFilterTest = []struct {
		tokens   []string
		expected []string
	}{
		{tokens: []string{"a", "and", "apple", "orange"}, expected: []string{"apple", "orange"}},
	}

	// TODO - FIXME
	proj.stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	for _, tt := range stopwordFilterTest {
		testname := fmt.Sprintf("%v", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			actual := proj.stopwordFilter(tt.tokens)
			if contains(actual, "an") || contains(actual, "and") {
				t.Errorf("got %v expected %v", actual, tt.expected)
			}
		})
	}
}

func TestStemmerFilter(t *testing.T) {
	var stemmerFilterTest = []struct {
		tokens   []string
		expected []string
	}{
		{tokens: []string{"fighting", "riding", "added", "ate"}, expected: []string{"fight", "ride", "add", "eat"}},
	}

	// TODO - FIXME
	proj.stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	for _, tt := range stemmerFilterTest {
		testname := fmt.Sprintf("%v", tt.tokens)
		t.Run(testname, func(t *testing.T) {
			actual := stemmerFilter(tt.tokens)
			if !contains(actual, "fight") || !contains(actual, "ride") {
				t.Errorf("got %v expected %v", actual, tt.expected)
			}
		})
	}
}
