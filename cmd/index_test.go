package main

import (
	"fmt"
	"reflect"
	"testing"
)

var p Project

func TestIntersection(t *testing.T) {
	var intersectionTestvals = []struct {
		a        []uint32
		b        []uint32
		expected []uint32
	}{
		{a: []uint32{1, 2, 3, 4, 5, 6, 7}, b: []uint32{1, 2, 5, 6}, expected: []uint32{1, 2, 5, 6}},
		{a: []uint32{3, 4, 5, 6}, b: []uint32{1, 2, 3, 5, 6}, expected: []uint32{3, 5, 6}},
	}

	for _, tt := range intersectionTestvals {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			res := intersection(tt.a, tt.b)
			t.Log("Intersection %n", res)
		})
	}
}

func TestAdd(t *testing.T) {
	var addDocumentToIndexTest = []struct {
		doc      document
		search   string
		expected []uint32
	}{
		// start with a simple document with 2 sentence
		// "today" repeated, should not add it again
		{doc: document{Id: 0, Text: "Today we are going to rIde. Hope today not gonna be raining"}, search: "Today", expected: []uint32{0}},
		// act "Today" again to get two items in the result
		{doc: document{Id: 1, Text: "Today all good."}, search: "Today", expected: []uint32{0, 1}},
		// just another doc
		{doc: document{Id: 2, Text: "It is ok to be not normal."}, search: "normal", expected: []uint32{2}},
		// intersection from the left
		{doc: document{Id: 3, Text: "Yet another document to search."}, search: "yet AnOther", expected: []uint32{3}},
		// intersection from the right
		{doc: document{Id: 3, Text: "This must be a longer sentence to have interesting result Today."}, search: "longer today", expected: []uint32{3}},
		// intersection using 3 terms
		{doc: document{Id: 4, Text: "Physics is exciting to study, even better if you like math too."}, search: "math physics study", expected: []uint32{4}},
		// add some trick to cover condition when res greater than Ids
		{doc: document{Id: 5, Text: "Physics is everywhere, it defines Today."}, search: "physics today", expected: []uint32{5}},
		// search term not in index
		{doc: document{Id: 6, Text: "A sentence which not contains the search term."}, search: "foo", expected: nil},
	}

	for _, tt := range addDocumentToIndexTest {
		testname := fmt.Sprintf("%v", tt.doc.Text)
		t.Run(testname, func(t *testing.T) {

			if len(p.idx) == 0 {
				p.idx = make(index)
			}

			p.documents = append(p.documents, tt.doc)
			p.add(tt.doc)
			t.Log("Index length: %n", len(p.idx))
			t.Log("Index content: %n", p.idx)

			sres := p.search(tt.search)
			for r := range sres {
				t.Logf("Search res %v", r)
			}

			d := p.documents[tt.doc.Id]
			t.Logf("Doc ID %v", d.Id)
			equality := reflect.DeepEqual(sres, tt.expected)
			if !equality {
				t.Errorf("Got %v, expected %v", sres, tt.expected)
			}
		})
	}
}

// TODO performance test
