package main

import "log"

type document struct {
	Id        uint32
	Key       string
	FilePath  string
	Title     string
	Text      string
	CreatedAt string
	UpdatedAt string
}

type documentProducer struct {
	docChan chan document
	quit    chan chan error
}

type index map[string][]uint32

type Project struct {
	files     []string
	idx       index
	query     string
	documents []document
	config    Config
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
}

type Config struct {
	Path         string `json:"path"`
	ShowContent  bool   `json:"show_content"`
	ShowNumLines uint16 `json:"show_num_lines"`
}
