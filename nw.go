package main

import (
	"os"
	"flag"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	wordPtr := flag.String("word", "", "A word to get alternatives to")
	flag.Parse()

	if *wordPtr == "" {
		fmt.Fprintf(os.Stderr, "No word passed\n")
		os.Exit(1)
	}

	url := "https://www.thesaurus.com/browse/" + *wordPtr

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error GETting " + url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error GETting " + url)
		os.Exit(1)
	}

	count := 0
	doc.Find(".synonyms-container").First().Find("ul > li").Each(func(i int, s *goquery.Selection) {
		word := s.Find("span").Text()
		length := len(word)
		spaceCount := 20 - length

		fmt.Printf(word)
		for i := 0; i < spaceCount; i++ {
			fmt.Printf(" ")
		}
		count++
		if count == 5 {
			count = 0
			fmt.Printf("\n")
		}
	})
	os.Exit(0)
}