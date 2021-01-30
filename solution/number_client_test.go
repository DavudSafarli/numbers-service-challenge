package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

type test struct {
	expected            []int
	serverResponse      []int
	serverSleepDuration time.Duration
	clientTimeout       time.Duration
}

func TestNumberClient(t *testing.T) {
	tests := []test{
		{
			expected:       []int{1, 2, 3, 4, 5},
			serverResponse: []int{1, 2, 3, 4, 5},
		},
		{
			expected:            []int{1, 2, 3, 4, 5},
			serverResponse:      []int{1, 2, 3, 4, 5},
			serverSleepDuration: 100 * time.Millisecond,
		},
		{
			expected:            []int{1, 2, 3, 4, 5},
			serverResponse:      []int{1, 2, 3, 4, 5},
			serverSleepDuration: 100 * time.Millisecond,
			clientTimeout:       10 * time.Millisecond,
		},
		{
			expected:            []int{1, 2, 3, 4, 5},
			serverResponse:      []int{1, 2, 3},
			serverSleepDuration: 100 * time.Millisecond,
			clientTimeout:       10 * time.Millisecond,
		},
		{
			expected:            []int{1, 2, 3, 4, 5},
			serverResponse:      []int{1, 2, 3, 4, 5},
			serverSleepDuration: 100 * time.Millisecond,
			clientTimeout:       200 * time.Millisecond,
		},
	}

	for i, tt := range tests {
		t.Run(`test NumberClient #`+strconv.Itoa(i), func(t *testing.T) {
			subject := NewHTTPNumberClient()
			testNumberClient(t, subject, tt)
		})
	}
}

func testNumberClient(t *testing.T, subject NumberClient, config test) {
	// start a server that sleeps `config.serverSleepDuration` and returns `config.serverResponse`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string][]int{
			"numbers": config.serverResponse,
		}
		if config.serverSleepDuration > 0 {
			time.Sleep(config.serverSleepDuration)
		}
		w.Header().Add("content-type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic("couldn't encode response in test")
		}
	}))

	// add WithTimeout if config.clientTimeout is specified
	ctx := context.Background()
	if config.clientTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, config.clientTimeout)
		// defer cancelFunc()
	}

	// act
	actual, err := subject.Get(ctx, server.URL)

	// if config.clientTimeout is specified, and server sleeps more than clientTimeout
	// then subject should respect the timeout and return error
	if config.clientTimeout != 0 && config.serverSleepDuration > config.clientTimeout {
		if err == nil {
			t.Error("Should have returned `context deadline exceeded error`.")
			t.FailNow()
		}
		return
	}

	// otherwise it shouldn't return error
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// and check actual/expected response
	if !isEq(actual, config.expected) {
		t.Error("fail\n",
			"actual   ", actual, "\n",
			"expected ", config.expected, "\n",
		)
		t.FailNow()
	}
}
