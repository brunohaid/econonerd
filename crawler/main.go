package crawler

import (
	"log"
	"time"
) 

// We keep them super simple for now
type Item struct {
	// To help with normalization we use the existing ids from reddit & twitter
	// and normalize all other urls before MD5 hashing them
	id				string			`json:"id"`
	kind			string			`json:"type"`
	title			string			`json:"title"`
	author			string			`json:"author"`
	published		time.Time		`json:"published"`
	firstseen		time.Time		`json:"firstseen"`
	url				string			`json:"url"`
	body			string			`json:"body"`
}

type Mention struct {
	// See id comment at Item
	id				string			`json:"id"`	
	target			string			`json:"target"`
	kind			string			`json:"type"`
	author			string			`json:"by"`
	url				string			`json:"url"`
	mentioned		time.Time 		`json:"timestamp"`
	title			string			`json:"title"`
	body			string			`json:"body"`
}

type Person struct {
	twitter			string			`json:"twitter"`
	reddit			string			`json:"reddit"`
	name 			string			`json:"name"`
	bio				string			`json:"bio"`
	urls			[]string 		`json:"urls"`	
}

type Alias struct {
	name 			string			`json:"name"`
	handle 			string			`json:"handle"`	
}

// Init
func Init() {
	// Log start
	log.Println("Spinning up crawler")

	// Kick off twitter crawling
	// crawltwitter()

	// Spawn RSS reader routines
	// crawlblogs()		

	// Spawn Subreddit subscriber routines
	crawlreddit()			
}