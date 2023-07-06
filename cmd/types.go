package main

type document struct {
	Title    string
	Text     string
	ID       int
	FileName string
	PathToFile string
}

type Project struct {
	dir       string
	files     []string
	idx       index
	query     string
	documents []document
}

type index map[string][]int