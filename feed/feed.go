package feed

import (
	"fmt"
	"log"
	"time"

	"github.com/hamza72x/blog-in-your-email/db"
	"github.com/hamza72x/blog-in-your-email/mail"
	"github.com/hamza72x/blog-in-your-email/model"
	"github.com/mmcdole/gofeed"
)

func CheckBlogFeed(blog model.Blog, isFirstRun bool) {

	log.Println("Fetching feed for: "+blog.Title, blog.Link)

	newPosts := getNewPosts(blog, isFirstRun)

	fmt.Printf("New posts of %s: %d\n", blog.Title, len(newPosts))

	for _, item := range newPosts {
		fmt.Printf("Sending email for %s, from %s\n", blog.Title, item.Author.Name)
		mail.Send(item.Content)
		time.Sleep(time.Second * 5)
	}
}

func getNewPosts(blog model.Blog, isFirstRun bool) []gofeed.Item {

	var newPosts []gofeed.Item

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(blog.Link)

	if err != nil {
		log.Println("Error parsing feed:", err)
		return []gofeed.Item{}
	}

	for i, item := range feed.Items {

		if i == 0 {
			newPosts = append(newPosts, *item)
		}

		var post model.Post

		db.Conn().Where("link = ?", item.Link).First(&post)

		if post.ID == 0 {

			db.Conn().Create(&model.Post{Link: item.Link, IsEmailSent: isFirstRun})

			if !isFirstRun {
				newPosts = append(newPosts, *item)
			}
		}
	}

	return newPosts
}
