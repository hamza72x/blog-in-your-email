package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func Send(email, message string) {

	// Sender data.
	from := os.Getenv("BLOG_IN_EMAIL")
	password := os.Getenv("BLOG_IN_PASSWORD")

	if len(from) == 0 {
		log.Printf("BLOG_IN_EMAIL is not set")
		return
	}

	if len(password) == 0 {
		log.Printf("BLOG_IN_PASSWORD is not set")
		return
	}

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}
