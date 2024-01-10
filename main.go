package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	query := searchCmd.String("q", "", "Search query")

	if len(os.Args) < 2 {
		fmt.Println("expected 'search' command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		HandleSearch(searchCmd, query)
	default:
		fmt.Println("unexpected command, type --help for commands")
		os.Exit(1)
	}

}

func HandleSearch(searchCmd *flag.FlagSet, query *string) {
	searchCmd.Parse(os.Args[2:])

	if *query == "" {
		fmt.Print("Search query is required")
		searchCmd.PrintDefaults()
		os.Exit(1)
	}

	search := *query

	fmt.Print(search)
}
