package main

import (
	"fmt"
	"github.com/gocolly/colly"
)
type Item struct {
	title string
	description string
	pubDate string
	link string
}

func main() {
	count := 0

	items := []Item{}

	c := colly.NewCollector()

	c.OnXML("//item", func(e *colly.XMLElement) {
		count++
		fmt.Println(count)
		item := Item{
			title: e.ChildText("title"),
			//description: e.ChildAttr("description","p"),
			pubDate: e.ChildText("pubDate"),
			link: e.ChildText("atom:link"),
		}
		items = append(items, item)
	})

	c.Visit("http://www.milliyet.com.tr/rss/rssNew/gundemRss.xml")

	for _, item := range items {
		fmt.Println("Title: ", item.title)
		//fmt.Println("Description: ", item.description)
		fmt.Println("PubDate: ", item.pubDate)
		fmt.Println("Link: ", item.link)
		fmt.Println()
	}
}