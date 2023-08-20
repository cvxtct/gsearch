package main

import (
	"bufio"
	"log"
	"os"
)

// TODO elaborate this function
// generic client to download,
// generic struct loader
func (p *Project) stopWords() {
	file, err := os.Open("english.txt")
	if err != nil {
		log.Panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	// Read file line by line.
	for fileScanner.Scan() {
		line := fileScanner.Text()
		p.stopwords[line] = struct{}{}
	}

	p.InfoLog.Printf("Stop words %d loaded successfully ", len(p.stopwords))
}
