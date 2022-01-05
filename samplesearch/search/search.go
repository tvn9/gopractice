package search

import (
	"log"
	"sync"
)

// map to store keys and values for matching search terms
var matchers = make(map[string]Matcher)

// Run
func Run(strm string) {
	// Retrieve feed and store in feeds
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// unbuffer chanel
	results := make(chan *Results)

	// Setup a waitgroup
	var waitgroup sync.WaitGroup

	// Set the number of goroutines
	waitgroup.Add(len(feeds))

	// Launch goroutine for each feed
	for _, fds := range feeds {

		// Retrieve a matcher for search
		matcher, exists := matchers[fds.Type]
		if !exists {
			matcher = matchers["Default"]
		}
		// launch goroutines to perform the seach
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, strm, results)
			waitgroup.Done()
		}(matcher, fds)
	}
	// Launch goroutines for all the network
	go func() {
		// Wait for everything to be processed
		waitgroup.Wait()

		// Close the channel to signal to the display
		close(results)
	}()
	// Start display results as they are available
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
