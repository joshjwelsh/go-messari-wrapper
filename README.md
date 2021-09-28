# Go Messari Web Api Wrapper

Description: This projects aims at providing a go web wrapper for the [Messari api](https://messari.io/api/docs).

The current version v1 only supports messari's standard api's. V2 plans to introduce authentication for entreprise api access. 

To see what api are available check out accessor.go.

## Query Examples

```
	field_query := accessor.FieldQuery{
 		ield: "?fields=id,slug,symbol,metrics/market_data/price_usd",
	}

	timeseries_query := accessor.TimeseriesQuery{
		Start:    "start=2021-04-15",
		End:      "end=2021-05-15",
		Interval: INTERVAL_1DY,
	}
```

## Example code 
```

func check(e error) {
	if e != nil {
		log.Fatalf("Main error: %v", e)
	}
}
func main() {
	accessor_ := accessor.NewAccessor()

	ts_query := accessor.TimeseriesQuery{
		Start:    "start=2021-04-15",
		End:      "end=2021-05-15",
		Interval: INTERVAL_1DY,
	}
	result, err := accessor_.GetPriceTimeseries("ETH", ts_query)()
	check(err)
	fmt.Println(result.Data.Schema)

}
```