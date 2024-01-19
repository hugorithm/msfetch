package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"
)

type item struct {
	Name       string
	Price      string
	ProductUrl string
}

func main() {
    searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
    query := searchCmd.String("q", "", "-q <search_query>")

    if len(os.Args) < 2 {
        fmt.Println("expected a command, type '--help' for commands.")
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
			Price:      h.ChildText("span.price-box span.final"),
			ProductUrl: h.ChildAttr("div.image-box a", "href"),
		}

		items = append(items, item)
	})

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Scrapping:", r.URL.String() + "\n")
    })

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error during website visit:", err)
	}

	printItemList(items)
}

func printItemList(items []item) {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Name", "Price", "Product URL"})
    table.SetRowLine(true)

    for _, i := range items {
        table.Append([]string{i.Name, i.Price, i.ProductUrl})
    }

    table.Render()
}
