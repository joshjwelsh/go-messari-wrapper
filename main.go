package main

import (
	"fmt"
	"log"

	"github.com/joshjwelsh/go-messari-wrapper/accessor"
)

var verbose bool = true

func check(e error) {
	if e != nil {
		log.Fatalf("Main error: %v", e)
	}
}
func main() {
	accessor := accessor.NewAccessor()

	// query := accessor.FieldQuery{
	// 	Field: "?fields=id,slug,symbol,metrics/market_data/price_usd",
	// }

	result, err := accessor.GetAsset("eth")()
	check(err)
	fmt.Println(result.Data)

}
