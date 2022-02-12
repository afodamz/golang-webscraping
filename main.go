package main

import (
	"strconv"
	"github.com/gocolly/colly"
	"fmt"
	"log"
	"encoding/json"
	"os"
)

// type Article struct {
// 	Title string `json:"title"`
// 	Description string `json:"description"`
// }

type Article struct {
	ID int `json:"id"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("starting webscraping")
	articles := make([]Article, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com/", "www.lindaikejisblog.com/"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {git remote add origin https://github.com/afodamz/golang-webscraping.git
		id,err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		desc := element.Text
	

		fact := Article{
			ID: id,
			Description: desc,
		}

		articles = append(articles, fact)

		collector.OnRequest(func(req *colly.Request){
			fmt.Println("Visiting", req.URL.String())
		})
		collector.Visit("https://www.factretriever.com/rhino-facts")

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent(""," ")
		enc.Encode(articles)
	})
}