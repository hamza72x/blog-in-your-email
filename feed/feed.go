package feed

import (
	"fmt"
	"log"

	"github.com/blog-in-your-email/data"
	"github.com/mmcdole/gofeed"
)

func GetFeed(blog data.Blog) {

	log.Println("Fetching feed for: "+blog.Title, blog.Link)

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(blog.Link)

	if err != nil {
		log.Println("Error parsing feed:", err)
		return
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
		fmt.Println(item.Content[0:100])
		fmt.Println(item.Link)
	}

	fmt.Println(feed.Title)
}
