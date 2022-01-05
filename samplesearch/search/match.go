package search

import "log"

// Results contains search results
type Results struct {
	Field   string
	Content string
}

// Matcher
type Matcher interface {
	Search(feed *Feed, strm string) ([]*Results, error)
}

// Match
func Match(matcher Matcher, feed *Feed, strm string, results chan<- *Results) {
	sResults, err := matcher.Search(feed, strm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range sResults {
		results <- result
	}
}

// Display
func Display(results chan *Results) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
