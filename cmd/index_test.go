package main

import (
	"fmt"
	"reflect"
	"testing"
)

var p Project

func TestAdd(t *testing.T) {
	var addDocumentToIndexTest = []struct {
		doc      document
		search   string
		expected []int
	}{	
		// start with a simple document with 2 sentence
		// "today" repeated, should not add it again
		{doc: document{ID: 0, Text: "Today we are going to ride. Hope today not gonna be raining"}, search: "Today", expected: []int{0}},
		// act "Today" again to get two items in the result
		{doc: document{ID: 1, Text: "Today all good."}, search: "Today", expected: []int{0, 1}},
		// just another doc
		{doc: document{ID: 2, Text: "It is ok to be not normal."}, search: "normal", expected: []int{2}},
		// intersection from the left
		{doc: document{ID: 3, Text: "Yet another document to search."}, search: "yet AnOther", expected: []int{3}},
		// intersection from the right
		{doc: document{ID: 3, Text: "This must be a longer sentence to have interesting result Today."}, search: "longer today", expected: []int{3}},
		// intersection using 3 terms
		{doc: document{ID: 4, Text: "Physics is exciting to study, even better if you like math too."}, search: "math physics study", expected: []int{4}},
		// add some trick to cover contition when res greater than ids
		{doc: document{ID: 5, Text: "Physics is everywhere, it defines Today."}, search: "physics today", expected: []int{5}},
		// search term not in index
		{doc: document{ID: 6, Text: "A sentence which not contains the search term."}, search: "foo", expected: nil},
	}

	for _, tt := range addDocumentToIndexTest {
		testname := fmt.Sprintf("%v", tt.doc.Text)
		t.Run(testname, func(t *testing.T) {

			if len(p.idx) == 0 {
				p.idx = make(index)
			}

			p.documents = append(p.documents, tt.doc)
			p.idx.add(tt.doc)
			t.Log("Index length: %n", len(p.idx))
			t.Log("Index content: %n", p.idx)

			sres := p.idx.search(tt.search)
			for r := range sres {
				t.Logf("Search res %v", r)
			}

			d := p.documents[tt.doc.ID]
			t.Logf("Doc ID %v", d.ID)
			equality := reflect.DeepEqual(sres, tt.expected)
			if !equality {
				t.Errorf("Got %v, expected %v", sres, tt.expected)
			}
		})
	}
}
