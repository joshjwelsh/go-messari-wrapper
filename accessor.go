package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Accessor struct {
	URL string
}

func NewAccessor() Accessor {
	return Accessor{
		URL: "https://data.messari.io/api",
	}
}

func handleQuery(api string, query ...Query) string {
	for _, q := range query {
		if q.Active() {
			api += q.Get()
		}
	}
	return api
}

func (a Accessor) GetAllAssets(query ...Query) func() (interface{}, error) {
	var api string = "/v2/assets"
	api = handleQuery(api, query...)
	if verbose {
		fmt.Println("api: ", api)
	}
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
			return response, fmt.Errorf("Error decoding response body: %v", err)
		}
		return response, nil
	}

}

func (a Accessor) GetAsset(asset string, query ...Query) func() (interface{}, error) {
	var api string = fmt.Sprintf("/v1/assets/%v", asset)
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
			return response, fmt.Errorf("Error decoding response body: %v", err)
		}
		return response, nil
	}
}

func (a Accessor) GetProfile(asset string, query ...Query) func() (interface{}, error) {
	var api string = fmt.Sprintf("/v2/assets/%v/profile", asset)
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
			return response, fmt.Errorf("Error decoding response body: %v", err)
		}
		return response, nil
	}

}
func (a Accessor) GetMetrics(asset string, query ...Query) func() (interface{}, error) {
	var api string = fmt.Sprintf("/v1/assets/%v/metrics", asset)
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
			return response, fmt.Errorf("Error decoding response body: %v", err)
		}
		return response, nil
	}
}

func (a Accessor) GetMarketData(asset string, query ...Query) func() (interface{}, error) {
	var api string = fmt.Sprintf("/v1/assets/%v/metrics/market-data", asset)
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
			return response, fmt.Errorf("Error decoding response body: %v", err)
		}
		return response, nil
	}
}

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
			return response, fmt.Errorf("Error decoding response body: %v", err)
		}
		return response, nil
	}
}
