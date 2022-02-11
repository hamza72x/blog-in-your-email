package main

import (
	"github.com/hamza72x/blog-in-your-email/data"
	"github.com/hamza72x/blog-in-your-email/db"
	"github.com/hamza72x/blog-in-your-email/feed"
)

func main() {

	blogs := data.GetBlogDataFromCSV()
	isFirstRun := db.IsFirstRun()

	for i := range blogs {
		feed.CheckBlogFeed(blogs[i], isFirstRun)
	}
}
