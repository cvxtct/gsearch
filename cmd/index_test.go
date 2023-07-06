package main

import (
	"fmt"
	"testing"
)


func TestAdd(t *testing.T) {
	var addDocumentToIndexTest = []struct {
		doc      document
		expected []int
	}{
		{doc: document{ID: 1, Text: "Lorem ipsum dolor. Sir amet ipsum."}},
		{doc: document{ID: 2, Text: "Lorem amestiter."}},
	}

	for _, tt := range addDocumentToIndexTest {
		testname := fmt.Sprintf("%v", tt.doc.ID)
		t.Run(testname, func(t *testing.T) {
			var p Project
			p.idx = make(index)
			
			p.documents = append(p.documents, tt.doc)
			p.idx.add(tt.doc)
			tt.doc.ID = 1
			res := p.idx.search("Lorem")
			t.Logf("sres %v", res)
			d := p.documents[0]
			t.Logf("Doc ID %v", d.ID)
			// FIXME
			// if d.ID != 1 {
			// 	t.Errorf("got %v", d.ID)
			// }
		})
	}
}
