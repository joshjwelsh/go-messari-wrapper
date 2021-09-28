package accessor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var verbose bool = true

type Accessor struct {
	URL string
}

func NewAccessor() Accessor {
	return Accessor{
		URL: "https://data.messari.io/api",
	}
}

// add queries to url
func handleQuery(api string, query ...Query) string {
	for _, q := range query {
		if q.Active() {
			api += q.Get()
		}
	}
	return api
}

func (a Accessor) GetAllAssets(query ...Query) func() (GetAllAssetsResponse, error) {
	var api string = "/v2/assets"
	api = handleQuery(api, query...)
	if verbose {
		fmt.Println("api: ", api)
	}
	return func() (GetAllAssetsResponse, error) {
		var response GetAllAssetsResponse
		url := a.URL + api
		res, err := http.Get(url)

		if err != nil {
			return response, fmt.Errorf("http.Get(%v) return an error: %v", url, err)
		}
		body := res.Body

		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)
		}
		return response, nil
	}

}

func (a Accessor) GetAsset(asset string, query ...Query) func() (GetAssetResponse, error) {
	var api string = fmt.Sprintf("/v1/assets/%v", asset)
	api = handleQuery(api, query...)
	return func() (GetAssetResponse, error) {
		var response GetAssetResponse
		url := a.URL + api
		res, err := http.Get(url)

		if err != nil {
			return response, fmt.Errorf("http.Get(%v) return an error: %v", url, err)
		}
		body := res.Body
		if verbose {
			log.Printf("Body in GetAllAssets: %v\n", body)
			log.Printf("Url used: %v\n", res.Request.URL)
			log.Printf("Status code received: %v\n", res.StatusCode)

		}
		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)
		}
		return response, nil
	}
}

func (a Accessor) GetProfile(asset string, query ...Query) func() (GetProfileResponse, error) {
	var api string = fmt.Sprintf("/v2/assets/%v/profile", asset)
	api = handleQuery(api, query...)
	return func() (GetProfileResponse, error) {
		var response GetProfileResponse
		url := a.URL + api
		res, err := http.Get(url)

		if err != nil {
			return response, fmt.Errorf("http.Get(%v) return an error: %v", url, err)
		}
		body := res.Body
		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)
		}
		return response, nil
	}

}
func (a Accessor) GetMetrics(asset string, query ...Query) func() (GetMetricsResponse, error) {
	var api string = fmt.Sprintf("/v1/assets/%v/metrics", asset)
	api = handleQuery(api, query...)
	return func() (GetMetricsResponse, error) {
		var response GetMetricsResponse
		url := a.URL + api
		res, err := http.Get(url)

		if err != nil {
			return response, fmt.Errorf("http.Get(%v) return an error: %v", url, err)
		}
		body := res.Body
		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)
		}
		return response, nil
	}
}

func (a Accessor) GetMarketData(asset string, query ...Query) func() (GetMarketDataResponse, error) {
	var api string = fmt.Sprintf("/v1/assets/%v/metrics/market-data", asset)
	api = handleQuery(api, query...)
	return func() (GetMarketDataResponse, error) {
		var response GetMarketDataResponse
		url := a.URL + api
		res, err := http.Get(url)

		if err != nil {
			return response, fmt.Errorf("http.Get(%v) return an error: %v", url, err)
		}
		body := res.Body
		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)
		}
		return response, nil
	}
}

// Currently there is not a detailed json to uses as reference so it returns an interface{}
func (a Accessor) GetAssetTimeseriesMetrics(asset string, query ...Query) func() (interface{}, error) {
	var api string = "/v1/assets/metrics"
	api = handleQuery(api, query...)
	return func() (interface{}, error) {
		var response interface{}
		url := a.URL + api
		res, err := http.Get(url)

		if err != nil {
			return response, fmt.Errorf("http.Get(%v) return an error: %v", url, err)
		}
		body := res.Body
		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)
		}
		return response, nil
	}
}

func (a Accessor) GetPriceTimeseries(asset string, query ...Query) func() (GetTimeseriesResponse, error) {
	var api string = fmt.Sprintf("/v1/assets/%v/metrics/price/time-series", asset)
	api = handleQuery(api, query...)
	return func() (GetTimeseriesResponse, error) {
		var response GetTimeseriesResponse
		url := a.URL + api
		res, err := http.Get(url)
		if err != nil {
			return response, fmt.Errorf("http.Get(%v) returned an error: %v", url, err)
		}
		body := res.Body
		defer body.Close()
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			return response, fmt.Errorf("error decoding response body: %v", err)

		}
		return response, nil
	}

}
