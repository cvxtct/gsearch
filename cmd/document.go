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
			fmt.Printf("Markdown file collected: %s\n", path)
		}
		return nil
	})
}

// parseDocument parsing .md file
func (p *Project) parseDocument(f string) (document, error) {
	var r document
	var lines string
	// Open file.
	file, err := os.Open(f)
	if err != nil {
		return document{}, err
	}

	fileScanner := bufio.NewScanner(file)

	// Read file line by line.
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines += line
	}
	
	r.Title = f
	r.Text = lines
	ss := strings.Split(f, "/")
	r.FileName = ss[len(ss)-1]
	ss = ss[:len(ss)-1]
	r.PathToFile = strings.Join(ss, "/")

	// Close file.
	file.Close()
	return r, nil
}
