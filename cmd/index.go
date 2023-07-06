package main


// it maps every word in documents to document IDs.
// the built-in map is a good candidate for storing the mapping.
// the key in the map is a token (string) and the value is a list of document IDs:
//func (idx index) add(docs []document){
func (idx index) add(doc document){
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
// and collect IDs that exist in both
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// updated search function to
func (idx index) search(text string) []int {
	var r []int
	for _, token := range normalize(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}
