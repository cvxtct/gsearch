package main

type document struct {
	Id        uint32
	Key       string
	FilePath  string
	Title     string
	Text      string
	CreatedAt string
	UpdatedAt string
}

type index map[string][]uint32

type Project struct {
	dir       string
	files     []string
	idx       index
	query     string
	documents []document
}
