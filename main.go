package main

import (
	"fmt"
	"github.com/gocolly/colly"
	cron2 "gopkg.in/robfig/cron.v2"
	"log"
	"strings"
)

func main() {
	runCron()
	fmt.Scanln()
}

func scrape() {
	c := colly.NewCollector()
	c.OnHTML("pre", func(e *colly.HTMLElement) {
		dataSlice := strings.Fields(e.Text)
		dataFiltered := dataSlice[42:53]
		date := dataFiltered[0]
		place := fmt.Sprintf("%s %s", dataFiltered[8], dataFiltered[9])
		intensity := dataFiltered[6]
		fmt.Printf("%s - %s - %s\n", date, place, intensity)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	c.Visit("http://www.koeri.boun.edu.tr/scripts/lst1.asp")
}

func runCron() {
	log.Println("Starting cron")
	cron := cron2.New()
	cron.AddFunc("@every 37s", scrape)

	cron.Start()
}
