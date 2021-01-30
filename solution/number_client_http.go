package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// HTTPNumberClient is an HTTP implementation of NumberClient
type HTTPNumberClient struct {
}

// NewHTTPNumberClient is a constructor for HTTPNumberClient
func NewHTTPNumberClient() HTTPNumberClient {
	return HTTPNumberClient{}
}

type numbersResponse struct {
	Numbers []int `json:"numbers"`
}

// Get makes an HTTP request to given `URL` and returns slice of ints
func (ms HTTPNumberClient) Get(ctx context.Context, URL string) ([]int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// map response body to numbersResponse struct
	value := numbersResponse{}
	err = json.NewDecoder(response.Body).Decode(&value)
	if err != nil {
		return nil, err
	}

	return value.Numbers, nil
}
