package summary

import (
	"fmt"
	"testing"
)

func TestTextRank(t *testing.T) {
	text := `Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language. It focuses on how to program computers to process and analyze large amounts of natural language data. The goal is to enable computers to understand human language in a valuable way.`

	sentences := splitSentences(text)
	summary := TextRank(sentences, 0.85, 20, 2) // dampingFactor=0.85, iterations=20, topN=2
	fmt.Println("TextRank Summary:", summary)
}
