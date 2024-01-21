package help 

import (
    "fmt"
    "os"
)

func HandleHelp() {
	fmt.Println("Commands:")
	fmt.Println("  search -q <search query>")
	fmt.Println("         -r <regional_indicator>")
	os.Exit(0)
}
