package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func (p *Project) interact() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// TODO make this bar dynamic
		color.Yellow("--------------------- s e a r c h ---------------------")
		color.Green("Documents indexed: %d", len(p.documents))
		color.Green("Index size: %d token(s)", len(p.idx))
		color.Blue(("Threads: %d\n"), threadProfile.Count())
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		p.query = text

		start := time.Now()
		matchedIDs := p.idx.search(p.query)
		elapsed := time.Since(start)

		if matchedIDs == nil {
			color.Red("Term %s not in index!", p.query)
			continue
		}

		log.Printf("Search found in %d document(s)", len(matchedIDs))

		for _, id := range matchedIDs {
			doc := p.documents[id]
			// TODO make this dynamic
			color.White("---------------------------------------------------")
			fmt.Printf("[In: %s]\t\nText: %s\n", doc.FilePath, doc.Text)
		}
		color.White("---------------------------------------------------")
		color.Yellow("Search took: %s", elapsed)
	}
}
