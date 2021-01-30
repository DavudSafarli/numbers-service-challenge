package main

import (
	"context"
	"sync"
	"time"
)

// NumberClient is a contract for retrieving numbers from a URL
type NumberClient interface {
	Get(ctx context.Context, URL string) ([]int, error)
}

// MergerSorter is a contact for merging and sorting slice of slices
type MergerSorter interface {
	MergeAndSort(ctx context.Context, slices [][]int) []int
}

// App is the business logic of application
type App struct {
	mergerSorter MergerSorter
	numberClient NumberClient
	waitDuration time.Duration
}

// NewApp is a constructor for App
func NewApp(mergerSorter MergerSorter, numberClient NumberClient, waitDuration time.Duration) App {
	return App{
		mergerSorter: mergerSorter,
		numberClient: numberClient,
		waitDuration: waitDuration,
	}
}

// Collect merges and sorts the integers returned
// by sending HTTP request to provided strings URLs
func (app App) Collect(ctx context.Context, urls []string) []int {
	collections := make([][]int, 0)
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	for _, URL := range urls {
		wg.Add(1)
		go func(URL string) {
			defer wg.Done()
			ctxtimeout, cancl := context.WithTimeout(context.Background(), app.waitDuration)
			defer cancl()
			numbers, _ := app.numberClient.Get(ctxtimeout, URL)

			if numbers == nil || len(numbers) == 0 {
				return
			}
			mx.Lock()
			defer mx.Unlock()
			collections = append(collections, numbers)
		}(URL)
	}
	wg.Wait()

	return app.mergerSorter.MergeAndSort(context.Background(), collections)
}
