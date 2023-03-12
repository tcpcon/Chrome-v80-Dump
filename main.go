package main

import (
	"dumper/chromium"

	"fmt"
)

func main() {
	for name, data := range chromium.Dump() {
		fmt.Printf("\n%s:\n", name)

		fmt.Printf("\t%d Accounts\n", len(data.Accounts))
		fmt.Printf("\t%d Cookies\n", len(data.Cookies))
		fmt.Printf("\t%d Web History\n", len(data.WebHistory))
		fmt.Printf("\t%d Search History\n", len(data.SearchHistory))
		fmt.Printf("\t%d Credit Cards\n", len(data.CreditCards))
		fmt.Printf("\t%d Autofill\n", len(data.Autofill))
	}
}
