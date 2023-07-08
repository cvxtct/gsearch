package main

import (
	"log"
	"runtime"
)

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

func main() {

	var p Project
	p.idx = make(index)

	log.Println("Starting Markdown search...")

	// checkers, check system before start
	// later check memory, check disk space, calculate! etc...
	log.Println("Pre flight checks...")
	preFlightCheckOs()
	log.Println("Pre flight checks done!")

	p.dir = "/Users/attilabalazs/Projects/__GO__"

	p.readFileNames()
	p.runIndex()

	// TODOs
	// cli program should store index on disk,
	// create file time hashes
	// recreate each file hash upon start,
	// if change -> reindex

	p.interact()
}
