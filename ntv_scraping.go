package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

// Entry represents a single RSS feed entry
type Entry struct {
	Title     string
	Subtitle  string
	ID        string
	Updated   string
	Link      string
}

func main() {
	// Instantiate the collector
	c := colly.NewCollector()

	// Array to hold the feed entries
	var entries []Entry

	// On every <entry> element which has child elements <title>, <id>, <updated>, <link>
	c.OnXML("feed/entry", func(e *colly.XMLElement) {
		//fmt.Println("Entry found")
		entry := Entry{
			Title:     e.ChildText("title"),
			ID:        e.ChildText("id"),
			Updated:   e.ChildText("updated"),
			Link:      e.ChildAttr("link", "href"),
		}
		entries = append(entries, entry)
	})

	// Start scraping the RSS feed
	err := c.Visit("https://www.ntv.com.tr/teknoloji.rss")
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, entry := range entries {
		fmt.Printf("Title: %s\nSubtitle: %s\nID: %s\nUpdated: %s\nLink: %s\n\n",
			entry.Title, entry.Subtitle, entry.ID, entry.Updated, entry.Link)
	}
}
