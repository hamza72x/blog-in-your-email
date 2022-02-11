package mail

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/hamza72x/blog-in-your-email/helper"
)

func Send(message string) {

	ini := helper.GetIni()

	from := ini.SENDER_EMAIL
	password := ini.SENDER_PASSWORD

	if len(from) == 0 {
		log.Printf("SENDER_EMAIl is not set")
		return
	}

	if len(password) == 0 {
		log.Printf("SENDER_PASSWORD is not set")
		return
	}

	to := []string{ini.RECEIVER_EMAIL}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}
