package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	link := "https://f95zone.to/threads/a-wifes-phone-v0-4-7-bloody-ink.153245/"

	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	line := doc.Find(".bbWrapper").First().Text()

	version, _ := regexp.Compile("Version\\: (?P<version>\\d{1,3}[\\.]{1}\\d{1,3}[\\.]{1}\\d{1,3})")
	updated, _ := regexp.Compile("Thread Updated\\: (?P<release>\\d{1,4}-{1}\\d{1,2}[-]{1}\\d{1,2})")
	fmt.Println(version.FindString(line))
	fmt.Println(updated.FindString(line))
}