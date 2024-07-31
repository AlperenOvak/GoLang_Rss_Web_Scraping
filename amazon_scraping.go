package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

type Item struct {
	ID    string `json:"id"`
	Class string `json:"class"`
	Type  string `json:"type"`
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time to execute %s: %v\n", name, time.Since(start))
	}
}

func main(){

	count := 0
 
	c := colly.NewCollector(colly.AllowedDomains("www.amazon.com.tr"))

	c.OnRequest(func(r *colly.Request){
			fmt.Println("Link of the page:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited %s with response code %d\n", r.Request.URL, r.StatusCode)
		//fmt.Println(string(r.Body)) // Print response body for inspection
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h*colly.HTMLElement){
		fmt.Println(count)
			/*h.ForEach("div.a-section.a-spacing-base", func(_ int, h*colly.HTMLElement){
					/*var name string
					name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
					var stars string
					stars = h.ChildText("span.a-icon-alt")
					var price string
					price = h.ChildText("span.a-price-whole")

					fmt.Println(count)
					count++

			})*/
	})

	c.Visit("https://www.amazon.com.tr/s?k=keyboard")
}
