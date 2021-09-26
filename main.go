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
	accessor_ := accessor.NewAccessor()

	// query := accessor.FieldQuery{

	// 	Field: "?fields=id,slug,symbol,metrics/market_data/price_usd",
	// }

	ts_query := accessor.TimeseriesQuery{
		Start:    "start=2021-04-15",
		End:      "end=2021-05-15",
		Interval: INTERVAL_1DY,
	}
	result, err := accessor_.GetPriceTimeseries("ETH", ts_query)()
	check(err)
	fmt.Println(result.Data.Schema)

}
