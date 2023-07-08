package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// readFileNames parses a directory after a specified file
func (p *Project) readFileNames() {

	log.Println("Scanning: ", p.dir)

	filepath.Walk(p.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if strings.Contains(path, ".md") {
			p.files = append(p.files, path)
			log.Printf("File collected: %s\n", path)
		}
		return nil
	})
}

// parseDocument parsing .md file
func (p *Project) parseDocument(f string) (*document, error) {
	//var doc document
	var lines string
	// Open file.
	file, err := os.Open(f)
	if err != nil {
		return &document{}, err
	}

	fileScanner := bufio.NewScanner(file)

	// Read file line by line.
	for fileScanner.Scan() {
		// TODO recognise title
		line := fileScanner.Text()
		lines += line + " "
	}

	// TODO create hash of document
	doc := document{
		Id:       uint32(len(p.documents)),
		Title:    f,
		Text:     lines,
		FilePath: f,
	}

	// Close file.
	file.Close()
	return &doc, nil
}
