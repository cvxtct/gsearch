package main

import "log"

// add function maps every word in documents to document IDs.
// the built-in map is a good candidate for storing the mapping.
// the key in the map is a token (string) and the value is a list of document IDs:
func (idx index) add(doc document){
	for _, token := range normalize(doc.Text) {
		ids := idx[token]
		log.Println("%n", ids)
		if ids != nil && ids[len(ids)-1] == doc.ID {
			// Don't add same ID twice.
			continue
		}
		idx[token] = append(ids, doc.ID)
	}
}

// intersection function iterates two lists simultaneously
// and collect IDs that are exist in both lists
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	res := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			res = append(res, a[i])
			i++
			j++
		}
	}
	return res
}

// search function checks serach term occurence
// in text using index and returns it
func (idx index) search(text string) []int {
	var res []int
	for _, token := range normalize(text) {
		if ids, ok := idx[token]; ok {
			// search term is one word
			if res == nil {
				res = ids
			// search term is more than one word
			} else {
				res = intersection(res, ids)
			}
		} else {
			// token doesn't exist.
			return nil
		}
	}
	return res
}
