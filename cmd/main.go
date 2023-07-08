package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)


func preFlightCheckOs() {
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		panic("Incompatible OS")
	case "darwin":
		log.Println("OS OSX OK!")
	case "linux":
		log.Println("OS Linux OK!")
	default:
		log.Printf("%s.\n", opsys)
	}
}

type producer struct {
	docChan chan document
	quit    chan chan error
}

func (p *producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {

	var start time.Time
	var elapsed time.Duration

	var p Project
	p.idx = make(index)

	prod := &producer{
		docChan: make(chan document, 10),
		quit:    make(chan chan error),
	}

	log.Println("Starting Markdown search...")

	// checkers, check system before start
	// later check memory, check disk space, calculate! etc...
	log.Println("Pre flight checks...")
	preFlightCheckOs()
	log.Println("Pre flight checks done!")

	p.dir = "/Users/attilabalazs/Projects/"

	p.readFileNames()
	fmt.Printf("LEN %d", len(p.files))
	// wg.Add(len(p.files))

	// producer
	go func() {
		for i, f := range p.files {
			wg.Add(1)
			doc := p.parseDocument(f, i)
		
			select {
			case prod.docChan <- doc:
			case ch := <-prod.quit:
				close(prod.docChan)
				// If the producer had an error while shutting down,
				// we could write the error to the ch channel here.
				close(ch)
				return
			}
		}
		defer close(prod.docChan)
		defer close(prod.quit)
	}()
	

	// consumer
	start = time.Now()
	for doc := range prod.docChan {
		if len(prod.docChan) > len(p.files) {
			err := prod.Close()
			fmt.Printf("unexpected error: %v\n", err)
		}
		// add to index
		p.idx.add(doc)
		// add to documents slice
		p.documents = append(p.documents, doc)
	}
	elapsed = time.Since(start)
	log.Printf("Indexing took: %s", elapsed)
	// TODOs
	// cli program should store index on disk,
	// create file time hashes
	// recreate each file hash upon start,
	// if change -> reindex

	reader := bufio.NewReader(os.Stdin)

	for {
		// TODO make this bar dynamic
		color.Yellow("--------------------- s e a r c h ---------------------")
		color.Green("Documents indexed: %d", len(p.documents))
		color.Green("Index size: %d token(s)", len(p.idx))
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		p.query = text

		start := time.Now()
		matchedIDs := p.idx.search(p.query)
		elapsed := time.Since(start)

		if matchedIDs == nil {
			color.Red("Term %s not in index!", p.query)
			continue
		}

		log.Printf("Search found in %d document(s)", len(matchedIDs))

		for _, id := range matchedIDs {
			doc := p.documents[id]
			// TODO make this dynamic
			color.White("---------------------------------------------------")
			fmt.Printf("[In: %s]\t\nText: %s\n", doc.FilePath, doc.Text)
		}
		color.White("---------------------------------------------------")
		color.Yellow("Search took: %s", elapsed)
	}

}
