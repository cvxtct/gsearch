package main

// add function maps every word in documents to document IDs.
// the key in the map is a token (string) and the value is a list of document IDs
// since the subesquent document id always greater
// document added to index by doc.ID results ascending index
func (idx index) add(doc *document) {
	for _, token := range normalize(doc.Text) {
		ids := idx[token]
		if ids != nil && ids[len(ids)-1] == doc.ID {
			// Don't add same ID twice.
			continue
		}
		idx[token] = append(ids, doc.ID)
	}
}

// intersection function iterates two lists simultaneously
// and collect IDs that are exist in both lists
// function do expect ascending indices!
func intersection(a []uint32, b []uint32) []uint32 {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	res := make([]uint32, 0, maxLen)
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

// search function retrieves document id(s) from index
func (idx index) search(text string) []uint32 {
	var res []uint32
	for _, token := range normalize(text) {
		if ids, ok := idx[token]; ok {
			// search term is one word
			if res == nil {
				res = ids
				// search term is more than one word
			} else {
				// intersection allows joined results
				res = intersection(res, ids)
			}
		} else {
			// token doesn't exist.
			return nil
		}
	}
	return res
}
