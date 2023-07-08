package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"runtime/pprof"

	"github.com/fatih/color"
)

// DocumentProcessor
func (p *Project) documentProcessor(f string) bool {
	// parse document
	doc, err := p.parseDocument(f)
	if err != nil {
		log.Fatal(err)
	}
	// append document to documents
	p.documents = append(p.documents, doc)
	// add document to index
	p.idx.add(doc)

	return true
}

// Runner is in charge to run Indexer in go routine
// TODO FIX parallel11ism!
func (p *Project) runner(boolChan chan bool) {
	for _, f := range p.files {
		indexed_file := p.documentProcessor(f)
		boolChan <- indexed_file
	}
}

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

var threadProfile = pprof.Lookup("threadcreate")

func main() {

	var p Project
	p.idx = make(index)

	log.Println("Starting Markdown search...")
	log.Printf(("Threads in starting: %d\n"), threadProfile.Count())

	// checkers, check system before start
	// later check memory, check disk space, calculate! etc...
	log.Println("Pre flight checks...")
	preFlightCheckOs()
	log.Println("Pre flight checks done!")

	// channel for sending back results of indexer
	boolChan := make(chan bool)
	defer close(boolChan)

	p.dir = "/Users/attilabalazs/Projects/"
	p.readFileNames()

	// TODOs
	// cli program should store index on disk,
	// create file time hashes
	// recreate each file hash upon start,
	// if change -> reindex
	start := time.Now()
	go p.runner(boolChan)
	// Wait for Runner to finish.
	for i := 0; i < len(p.files); i++ {
		res := <-boolChan
		if res {
			log.Println("Item: ", p.files[i], " indexed!")
		}
		if i == len(p.files)-1 {
			log.Println("All files are indexed!")
			break
		}
	}


	elapsed := time.Since(start)
	log.Printf("Indexing took: %s", elapsed)

	reader := bufio.NewReader(os.Stdin)

	for {
		// TODO make this bar dynamic
		color.Yellow("--------------------- s e a r c h ---------------------")
		color.Green("Documents indexed: %d", len(p.documents))
		color.Green("Index size: %d token(s)", len(p.idx))
		color.Blue(("Threads: %d\n"), threadProfile.Count())
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		p.query = text

		start = time.Now()
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
