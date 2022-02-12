package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/hamza72x/blog-in-your-email/helper"
	"github.com/hamza72x/blog-in-your-email/tmpl"
	hel "github.com/hamza72x/go-helper"
	"github.com/mmcdole/gofeed"
	mail "github.com/xhit/go-simple-mail/v2"
)

func Send(item *gofeed.Item, feedTitle string) {
	send(fmt.Sprintf("New Article: %s", item.Title), tmpl.GetHtml(item, feedTitle))
}

func SendWelcomeEmail() {
	send("Welcome to BLOG IN YOUR EMAIL", hel.FileStrMust("tmpl/welcome.html"))
}

func send(subject, body string) {

	ini := helper.GetIni()

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = ini.SMTP_SERVER
	server.Port = ini.SMTP_PORT
	server.Username = ini.SENDER_EMAIL
	server.Password = ini.SENDER_PASSWORD
	server.Encryption = mail.EncryptionSTARTTLS

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Printf("Error connecting to SMTP server: %s", err)
		return
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()

	email.SetFrom(fmt.Sprintf("%s <%s>", "BLOG IN YOUR EMAIL", ini.SENDER_EMAIL))
	email.AddTo(ini.RECEIVER_EMAIL)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	// always check error after send
	if email.Error != nil {
		log.Printf("Error setting data to email: %s", email.Error)
		return
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(fmt.Sprintf("Email sent %s", subject))
	}
}
