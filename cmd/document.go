package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// readFileNames parses a directory after a specified file
func (p *Project) readFileNames() {

	log.Println("Scanning: ", p.config.Path)

	err := filepath.Walk(p.config.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			p.ErrorLog.Fatalf(err.Error())
		}
		if strings.Contains(path, ".md") {
			p.files = append(p.files, path)
			p.InfoLog.Printf("File collected: %s\n", path)
		}
		return nil
	})
	
	if err != nil {
		panic(err)
	}
}

// parseDocument parsing .md file
func (p *Project) parseDocument(f string, i int) document {
	defer wg.Done()
	//var doc document
	var lines string
	// Open file.
	file, err := os.Open(f)
	if err != nil {
		return document{}
	}

	fileScanner := bufio.NewScanner(file)

	// Read file line by line.
	for fileScanner.Scan() {
		// TODO recognize title
		line := fileScanner.Text()
		lines += line + " "
	}

	// TODO create hash of document
	doc := document{
		Id:       uint32(i),
		Title:    f,
		Text:     lines,
		FilePath: f,
	}

	// Close file.
	file.Close()
	p.InfoLog.Printf("Document %s parsed!", f)
	return doc
}
