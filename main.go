package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type Item struct {
	Price string
	Date  string
	Link  string
	Text  string
}

func main() {

	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	headers := []string{"Price", "Date", "Link", "Text"}
	writer.Write(headers)

	c := colly.NewCollector(
		colly.AllowedDomains("www.dba.dk"),
	)

	c.OnHTML(".dbaListing.listing", func(e *colly.HTMLElement) {
		regexp := regexp.MustCompile(`[^0-9]`)
		item := Item{}

		item.Price = regexp.ReplaceAllString(e.ChildText("[title='Pris']"), "")
		item.Date = e.ChildText("[title='Dato']")
		item.Link = e.ChildAttr("a", "href")

		replacer := strings.NewReplacer("-", "", ",", " ", "\"", "", "\n", "")
		item.Text = replacer.Replace(e.ChildText(".listingLink"))

		row := []string{item.Price, item.Date, item.Link, item.Text}
		writer.Write(row)
	})

	c.OnHTML(".trackClicks.a-page-link", func(e *colly.HTMLElement) {
		nextpage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextpage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitng", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println(response.StatusCode)
	})

	c.Visit("https://www.dba.dk/soeg/?soeg=gtx+1080+ti")

	//sortCSV()
}

// func sortCSV() {
// 	csvfile, err := os.Open("export.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	df := dataframe.ReadCSV(csvfile)

// 	sorted := df.Arrange(dataframe.Sort("Price"))

// 	f, err := os.Create("1export.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sorted.WriteCSV(f)
// }
