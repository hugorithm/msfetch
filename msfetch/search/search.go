package search 

import (
    "flag"
    "fmt"
    "os"
    "strings"

    "msfetch/scrape"
)

func HandleSearch(searchCmd *flag.FlagSet, query *string, region *string) {
    var defaultRegion string = "en_CH/CHF/"

    searchCmd.Parse(os.Args[2:])

    if *query == "" {
        fmt.Print("Search query is required\n\n")
        searchCmd.PrintDefaults()
        os.Exit(1)
    }

    if *region != "" {
       defaultRegion = HandleRegion(searchCmd, region)
    }
    
    scrape.Scrape(query, &defaultRegion)
}

func HandleRegion(searchCmd *flag.FlagSet, region *string) string {
    searchCmd.Parse(os.Args[2:])

    if *region == "" {
        fmt.Print("Regional indicator is required \n\n")
        searchCmd.PrintDefaults()
        os.Exit(1)
    }

    switch strings.ToLower(*region) {
    case "de":
        return "en_DE/EUR/"
    case "pt":
        return "en_PT/EUR/"
    case "uk":
        return  "en_GB/GBP/"
    case "us":
        return "en_US/USD/"
    case "ch":
        return "en_CH/CHF/"
    case "es":
        return "en_ES/EUR/"
    default:
        fmt.Println("Unexpected regional indicator. Using default \"CH\".")
        return "en_CH/CHF/"
    }
}
