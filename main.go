package main

import (
	"github.com/hamza72x/blog-in-your-email/data"
	"github.com/hamza72x/blog-in-your-email/feed"
)

func main() {

	checkConfigs()

	blogs := data.GetBlogDataFromCSV()

	for i := range blogs {
		feed.CheckBlogFeed(blogs[i])
	}
}

func checkConfigs() {

}
