package main

type document struct {
	Title      string
	Text       string
	ID         uint32
	FileName   string
	PathToFile string
}

type index map[string][]uint32

type Project struct {
	dir       string
	files     []string
	idx       index
	query     string
	documents []*document
}
