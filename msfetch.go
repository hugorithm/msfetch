package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name       string
	Price      string
	ProductUrl string
}

func main() {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	query := searchCmd.String("q", "", "Search query")

	if len(os.Args) < 2 {
		fmt.Println("expected a command, type '--help' to see available commands.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		HandleSearch(searchCmd, query)
	case "--help":
		HandleHelp()
	default:
		fmt.Println("unexpected command, type '--help' for commands")
		os.Exit(1)
	}
}

func Scrape(query *string) {
	escapedQuery := url.QueryEscape(*query)
	url := "https://www.musicstore.com/en_CH/CHF/search?SearchTerm=" + escapedQuery

	c := colly.NewCollector(
		colly.AllowedDomains("www.musicstore.com"),
	)

	var items []item

	c.OnHTML("div.tile-product", func(h *colly.HTMLElement) {
		item := item{
			Name:       h.ChildText("div[data-dynamic-block-name=ProductTile-ProductTitle] a span"),
			Price:      h.ChildText(".price-box span"),
			ProductUrl: h.ChildAttr("div.image-box a", "href"),
		}

		items = append(items, item)
	})

	c.OnHTML("div.settings-box-top a[title='to next page']", func(h *colly.HTMLElement) {
		c.Visit(h.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error during website visit:", err)
	}
	fmt.Println(items)
}

func HandleSearch(searchCmd *flag.FlagSet, query *string) {
	searchCmd.Parse(os.Args[2:])

	if *query == "" {
		fmt.Print("Search query is required\n\n")
		searchCmd.PrintDefaults()
		os.Exit(1)
	}

	Scrape(query)
}

func HandleHelp() {
	fmt.Println("Commands:")
	fmt.Println("  search -q <search query>")
	os.Exit(0)
}
