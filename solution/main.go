package main

import "time"

func main() {
	mergersorter := NewDefaultMergerSorter()
	numberClient := NewHTTPNumberClient()
	waitDuration := 500 * time.Millisecond
	app := NewApp(mergersorter, numberClient, waitDuration)

	server(app)
}
