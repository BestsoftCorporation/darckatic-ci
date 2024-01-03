package main

import (
	"darkatic-ci/cmd"
	"darkatic-ci/gui"
	"darkatic-ci/internal/api"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Increment the wait group counter for each goroutine
	wg.Add(3)

	// Start goroutines
	go func() {
		defer wg.Done()
		api.StartServer()
	}()

	go func() {
		defer wg.Done()
		gui.ServeGUI()
	}()

	go func() {
		defer wg.Done()
		cmd.Execute()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
