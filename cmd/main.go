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

func (p *Project) preFlightCheckOs() {
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		panic("Incompatible OS")
	case "darwin":
		p.InfoLog.Println("OS OSX OK!")
	case "linux":
		p.InfoLog.Println("OS Linux OK!")
	default:
		p.InfoLog.Printf("%s.\n", opsys)
	}
}

func (d *documentProducer) Close() error {
	ch := make(chan error)
	d.quit <- ch
	return <-ch
}

func main() {

	// performance times
	var start time.Time
	var elapsed time.Duration

	// init project
	var p Project
	p.idx = make(index)

	// read config
	p.config = Configuration()
	// create loggers
	p.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	p.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	docProd := &documentProducer{
		docChan: make(chan document, 100),
		quit:    make(chan chan error),
	}

	p.InfoLog.Println("Starting Markdown search...")

	// checkers, check system before start
	// later check memory, check disk space, calculate! etc...
	p.InfoLog.Println("Pre flight checks...")
	p.preFlightCheckOs()
	p.InfoLog.Println("Pre flight checks done!")

	// read file names from the path given in config
	p.readFileNames()

	// document producer
	go func() {
		for i, f := range p.files {
			wg.Add(1)
			doc := p.parseDocument(f, i)

			select {
			case docProd.docChan <- doc:
			case ch := <-docProd.quit:
				close(docProd.docChan)
				// If the producer had an error while shutting down,
				// we could write the error to the ch channel here.
				close(ch)
				return
			}
		}
		defer close(docProd.docChan)
		defer close(docProd.quit)
	}()

	// document consumer
	start = time.Now()
	for doc := range docProd.docChan {
		if len(docProd.docChan) > len(p.files) {
			err := docProd.Close()
			fmt.Printf("unexpected error: %v\n", err)
		}
		// add to index
		p.add(doc)
		// add to documents slice
		p.documents = append(p.documents, doc)
	}
	elapsed = time.Since(start)
	p.InfoLog.Printf("Indexing took: %s", elapsed)
	// TODOs
	// cli program should store index on disk,
	// create file time hashes
	// recreate each file hash upon start,
	// if change -> reindex

	reader := bufio.NewReader(os.Stdin)

	// User interaction

	for {
		fmt.Print("\n")
		color.Yellow("//////////////////////////////////////////// s e a r c h /////////////////////////////////////////")
		fmt.Print("\n")
		color.Green("Documents indexed: %d", len(p.documents))
		color.Green("Index size: %d token(s)", len(p.idx))
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		p.query = text

		start := time.Now()
		matchedIDs := p.search(p.query)
		elapsed := time.Since(start)

		if matchedIDs == nil {
			color.Red("Term << %s >> not in index!", p.query)
			continue
		}

		fmt.Print("\n")
		color.Yellow("Search found in %d document(s)", len(matchedIDs))
		fmt.Print("\n")

		// Printing out results

		// find longest path
		maxLen := 0
		for _, id := range matchedIDs {
			doc := p.documents[id]
			if maxLen < len(doc.FilePath) {
				maxLen = len(doc.FilePath)
			}
		}

		// ruler width
		var i int
		ruler := "------"

		if maxLen < 91 {
			maxLen = 90
		}

		for i < maxLen {
			ruler += "-"
			i++
		}

		// create formated result
		for i, id := range matchedIDs {
			doc := p.documents[id]

			color.White("%s", ruler)
			color.White("[%v][%s]\t\n", i, doc.FilePath)

			// align text to ruler if set in config
			if p.config.ShowContent {

				color.White("%s", ruler)

				formated := ""
				var lineCount uint16
				lineCount = 0
				for i := 0; i < len(doc.Text); i++ {
					// do not print more lines then set in
					if lineCount == p.config.ShowNumLines {
						break
					}
					formated += string(doc.Text[i])
					if i%(maxLen+6) == 0 && i != 0 {
						formated += "\n"
						lineCount += 1
					}
				}

				color.HiBlack("%s", formated)
				color.White("[max lines: %d]\n", p.config.ShowNumLines)
				color.White("[hash: %x]", doc.Key)
				color.White("[created: %v]\n", doc.CreatedAt.UTC())
				color.White("%s", ruler)
				fmt.Print("\n")
			}
		}
		fmt.Print("\n")
		color.Yellow("Search took: %s\n", elapsed)
		fmt.Print("\n")
	}
}
