package main

import (
	"log"
	"time"
)

func (p *Project) runIndex() {
	var start time.Time
	var elapsed time.Duration

	for _, f := range p.files {
		doc, err := p.parseDocument(f)
		if err != nil {
			log.Fatal(err)
		}
		// append document to documents
		p.documents = append(p.documents, doc)

		start = time.Now()
		p.idx.add(doc)
		elapsed = time.Since(start)
	}

	log.Printf("Indexing took: %s", elapsed)
}
