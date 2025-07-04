package ppt_crawling

import (
	"fmt"
	"log"
	"net/smtp"
)

func commandMail(cfg *Config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("usage: mail")
	}

	// Sender data.
	from := "youremail@gmail.com"
	password := "your-app-password" // NOT your Gmail password!

	// Receiver email address.
	to := []string{
		"receiver@example.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Subject: Test Email from Go\n\nThis is the email body!")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email Sent Successfully!")
	return nil
}
