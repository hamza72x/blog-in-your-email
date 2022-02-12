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

	t, err = t.Parse(EMAIL_HTML)

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

const EMAIL_HTML = `
<!doctype html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.FeedTitle}}</title>
    <style>
        body {
            width: 90% !important;
            page-break-after: always;
        }

        body *,
        body a,
        body p,
        body div {
            font-family: "Noto Sans Bengali", serif !important;
            color: #000 !important;
            margin-block-start: 0 !important;
            margin-block-end: 0 !important;
            font-weight: 400;
            font-size: 18px !important;
        }

        body {
            background: #FFF !important;
        }

        body p,
        body {
            text-align: justify;
        }

        body p {
            font-size: 18px !important;
            font-weight: 400;
        }

        body h1,
        body h2,
        body h3,
        body h4 {
            font-weight: bolder !important;
        }

        body h1 {
            font-size: 21px !important;
        }

        body h2 {
            font-size: 20px !important;
        }

        body h3 {
            font-size: 19px !important;
        }

        body h4 {
            font-size: 18px !important;
        }

        body h5 {
            font-size: 17px !important;
        }

        body h6 {
            font-size: 12px !important;
        }

        body h1 a,
        body h3 a,
        body h4 a {
            font-weight: bolder !important;
        }

        img,
        iframe {
            max-width: 95% !important;
            object-fit: contain;
            margin: 0 auto !important;
            height: auto;
            -webkit-box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
            -moz-box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
            box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
            -webkit-border-radius: 10px;
            -moz-border-radius: 10px;
            border-radius: 10px;
        }

        .entry-header {
            text-align: center;
        }
    </style>
</head>

<body>
    <h2>{{.Item.Title}}</h2>
    <br>
    {{.Content}}
    <br>
    <hr>
    <br>
    <a href="{{.Item.Link}}">{{.Item.Title}}</a>
</body>

</html>
`

const WELCOME_HTML = `
<!doctype html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Welcome email from blog-in-your-email</title>
    <style>
        body {
            width: 90% !important;
            page-break-after: always;
        }

        body *,
        body a,
        body p,
        body div {
            font-family: "Noto Sans Bengali", serif !important;
            color: #000 !important;
            margin-block-start: 0 !important;
            margin-block-end: 0 !important;
            font-weight: 400;
            font-size: 18px !important;
        }

        body {
            background: #FFF !important;
        }

        body p,
        body {
            text-align: justify;
        }

        body p {
            font-size: 18px !important;
            font-weight: 400;
        }

        body h1,
        body h2,
        body h3,
        body h4 {
            font-weight: bolder !important;
        }

        body h1 {
            font-size: 21px !important;
        }

        body h2 {
            font-size: 20px !important;
        }

        body h3 {
            font-size: 19px !important;
        }

        body h4 {
            font-size: 18px !important;
        }

        body h5 {
            font-size: 17px !important;
        }

        body h6 {
            font-size: 12px !important;
        }

        body h1 a,
        body h3 a,
        body h4 a {
            font-weight: bolder !important;
        }

        img,
        iframe {
            max-width: 95% !important;
            object-fit: contain;
            margin: 0 auto !important;
            height: auto;
            -webkit-box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
            -moz-box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
            box-shadow: 0 0 3px 0 rgba(0, 0, 0, .12), 0 1px 2px rgba(0, 0, 0, .24);
            -webkit-border-radius: 10px;
            -moz-border-radius: 10px;
            border-radius: 10px;
        }

        .entry-header {
            text-align: center;
        }
    </style>
</head>

<body>
    <h2>This is the first email from blog-in-your-email, just to make sure that email configuration is working correctly!
    </h2>
    <br>
    <hr>
    <a href="https://github.com/hamza72x/blog-in-your-email">blog-in-your-email</a>
</body>

</html>
`
