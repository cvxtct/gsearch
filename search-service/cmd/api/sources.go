package main

import (
	"bufio"
	"log"
	"os"
)

func (p *Project) ReadSourceFile(source string) {
	file, err := os.Open(source)
	if err != nil {
		log.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	// Read file line by line.
	for fileScanner.Scan() {
		line := fileScanner.Text()
		p.stopWords = map[string]struct{}{line: {}}
	}
}
