package main

import (
	"flag"
	"fmt"
	"os"

	"msfetch/help"
	"msfetch/search"
)

func main() {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	query := searchCmd.String("q", "", "-q <search_query>")
	region := searchCmd.String("r", "", "-r <regional_indicator>")

	if len(os.Args) < 2 {
		fmt.Println("Expected a command, type '--help' for commands.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		search.HandleSearch(searchCmd, query, region)
	case "--help":
		help.HandleHelp()
	default:
		fmt.Println("Unexpected command, type '--help' for commands")
		os.Exit(1)
	}
}
