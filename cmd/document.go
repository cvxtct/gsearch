package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ParseFiles parses a directory.
func (p *Project) readFileNames() {

	log.Println("Scanning: ", p.dir)

	filepath.Walk(p.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if strings.Contains(path, ".md") {
			p.files = append(p.files, path)
			fmt.Printf("File collected: %s\n", path)
		}
		return nil
	})
}

// parseDocument parsing .md file
func (p *Project) parseDocument(f string) (*document, error) {
	var doc document
	var lines string
	// Open file.
	file, err := os.Open(f)
	if err != nil {
		return &document{}, err
	}

	fileScanner := bufio.NewScanner(file)

	// Read file line by line.
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines += line + " "
	}

	// TODO create hash of document
	// TODO use documents length to count document id
	doc.ID = uint32(len(p.documents))
	doc.Title = f // file name with path // TODO replace with something else
	doc.Text = lines // lines extracted from doc
	ss := strings.Split(f, "/")
	doc.FileName = ss[len(ss)-1] // file name
	ss = ss[:len(ss)-1]
	doc.PathToFile = strings.Join(ss, "/") // path to file

	// Close file.
	file.Close()
	return &doc, nil
}
