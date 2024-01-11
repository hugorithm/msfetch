package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name        string
	Description string
	Price       string
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
	url := "https://www.musicstore.com/en_PT/EUR/search?SearchTerm=" + escapedQuery

	c := colly.NewCollector(
		colly.AllowedDomains("www.thomann.de", "www.musicstore.com"),
	)

  var items []item

	c.OnHTML("div.product", func(h *colly.HTMLElement) {
    fmt.Println(h.ChildText("span.product__price-primary"))
    item := item {
      Name: h.ChildText(".title__manufacturer") + h.ChildText(".title__name"),
      Description: h.ChildText(".product__description-item"),
      Price: h.ChildText(".product__price-primary"),
    }

    items = append(items, item)
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
