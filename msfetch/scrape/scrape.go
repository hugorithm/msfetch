package scrape 

import (
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

func Scrape(query *string, defaultRegion *string) {
	escapedQuery := url.QueryEscape(*query)
	url := "https://www.musicstore.com/" + *defaultRegion + "search?SearchTerm=" + escapedQuery

    c := colly.NewCollector(
		colly.AllowedDomains("www.musicstore.com", "www.musicstore.de", "www.dv247.com"),
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

    if len(items) <= 0 {
        fmt.Println("No resulst were found that match the serach query.")
        return
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
