package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
	"log"
	"encoding/json"
)

type Item struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Price string `json:"price"`
}


func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time to execute %s: %v\n", name, time.Since(start))
	}
} 	

func main() {
	defer timer("scraping")()
	c := colly.NewCollector(colly.Async(true))

	items := []Item{}

	c.OnHTML("div.side_categories li ul li", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		c.Visit(h.Request.AbsoluteURL(link))
	})

	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		i := Item{
			Title: h.ChildText("h3"),
			Link:  h.ChildAttr("h3 a", "href"),
			Price: h.ChildText("p.price_color"),
		}
		items = append(items, i)
	})


	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://books.toscrape.com/")
	c.Wait()

	data, err := json.MarshalIndent(items,"","  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}