package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("#postsList", func(e *colly.HTMLElement) {
		e.ForEach(".text-text-title", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("h2"),
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://haraj.com.sa/search/شريك?duringdate=1months")
}
