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

// Indexer
func (p *Project) Indexer(f string, i int) bool {
	doc, err := p.loadMarkDownDocuments(f)
	if err != nil {
		log.Fatal(err)
	}
	doc.ID = i
	p.documents = append(p.documents, doc)
	p.idx.add(doc)

	return true
}

// Runner is in charge to run Indexer in go routine
func (p *Project) Runner(boolChan chan bool) {
	for i, f := range p.files {
		indexed_file := p.Indexer(f, i)
		boolChan <- indexed_file
	}
}

func checkOs() {
	opsys := runtime.GOOS
	switch opsys {
	case "windows":
		panic("Incompatible OS")
	case "darwin":
		log.Println("MAC operating system")
	case "linux":
		log.Println("Linux")
	default:
		log.Printf("%s.\n", opsys)
	}
}

func main() {
	var p Project
	p.idx = make(index)

	log.Println("Starting Markdown search...")

	// checkers, check system before start
	// later check memory, check disk space, calculate! etc...
	checkOs()

	// channel for sending back results of indexer
	boolChan := make(chan bool)
	defer close(boolChan)

	//p.ParseArgument()
	p.dir = "/Users/attilabalazs/Projects/__GO__"
	p.FindFiles()

	// TODOs
	// move indexer outside from main, do not index if index exists
	// cli program should store index on disk,
	// create file time hashes
	// recreate each file hash upon start,
	// if change -> reindex
	start := time.Now()
	go p.Runner(boolChan)
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

	// mux := http.NewServeMux()
	// mux.HandleFunc("/custom_debug_path/profile", pprof.Profile)
	// log.Fatal(http.ListenAndServe(":7777", mux))

	for {
		// TODO make this dynamic
		color.Yellow("--------------------- s e a r c h ---------------------")
		color.Green("Documents indexed: %d", len(p.documents))
		color.Green("Index size: %d", len(p.idx))
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		p.query = text

		start = time.Now()
		matchedIDs := p.idx.search(p.query)
		elapsed := time.Since(start)
		log.Printf("Search found in %d document(s)", len(matchedIDs))

		for _, id := range matchedIDs {
			doc := p.documents[id]
			// TODO make this dynamic
			color.White("--------------------------------------------------------------------------------------------")
			fmt.Printf("[In: %s/%s]\tContent: %s\n", doc.PathToFile, doc.FileName, doc.Text)
		}
		log.Printf("Search took: %s", elapsed)
	}
}
