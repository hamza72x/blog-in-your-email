package tmpl

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/mmcdole/gofeed"
)

func GetHtml(item *gofeed.Item, feedTitle string) string {

	t := template.New("email.html")
	var err error

	t, err = t.ParseFiles("tmpl/email.html")

	if err != nil {
		log.Printf("Error parsing email template: %s\n", err)
		return item.Content
	}

	data := struct {
		FeedTitle string
		Item      gofeed.Item
		Content   template.HTML
	}{
		FeedTitle: feedTitle,
		Item:      *item,
		Content:   template.HTML(item.Content),
	}

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, data); err != nil {
		fmt.Printf("Error executing email template: %s\n", err)
		return item.Content
	}

	return tpl.String()
}
