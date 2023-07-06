package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// ParseArgument parses project path from command line.
// func (p *Project) ParseArgument() {
// 	flag.StringVar(&p.dir, "dir", "", "Directory to scan")
// 	flag.StringVar(&p.query, "q", "Test term", "search query")
// 	flag.Parse()
// }

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

	log.Println("Starting program")

	// checkers, check system before start
	// later check memory, check disk space, calculate! etc...
	checkOs()
	
	// Channel for sending back results of .
	boolChan := make(chan bool)
	defer close(boolChan)

	//p.ParseArgument()
	p.dir = "/Users/attilabalazs/Projects/__ACTIVE__/"
	p.FindFiles()
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

	// fmt.Print(p.idx)

	reader := bufio.NewReader(os.Stdin)

	// TODO make this dynamic
	fmt.Println("--------------------- s e a r c h ---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		p.query = text

		start = time.Now()
		matchedIDs := p.idx.search(p.query)
		log.Printf("Search found in %d document(s)", len(matchedIDs))
		elapsed := time.Since(start)

		for _, id := range matchedIDs {
			doc := p.documents[id]
			// TODO make this dynamic
			fmt.Println("--------------------------------------------------------------------------------------------")
			fmt.Printf("[In: %s/%s]\tContent: %s\n", doc.PathToFile, doc.FileName, doc.Text)
		}
		log.Printf("Search took: %s", elapsed)
		log.Printf("Documents indexed %d", len(p.documents))
		log.Printf("Index size %d", len(p.idx))
	}
}
