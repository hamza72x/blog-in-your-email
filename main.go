package main

import (
	"github.com/blog-in-your-email/data"
	"github.com/blog-in-your-email/feed"
)

func main() {
	// mail.Send("hamza@gtaf.org", "Hello World!")
	blogs := data.GetBlogDataFromCSV()
	for _, blog := range blogs {
		feed.GetFeed(blog)
	}
}

func makeFormattedMailMessage() {
	// Formatted message.

}
