package net

import (
	"log"
	"net/smtp"
	"os"
	"testing"
)

func TestSendMailSimple(t *testing.T) {
	from := "support@x.com"
	user := "support@x.com"
	pass := os.Getenv("MAIL_PASS")
	to := []string{
		"jdoe@qmail.com",
	}
	addr := "smtp.porkbun.com:587"
	host := "smtp.porkbun.com"

	msg := []byte("From: john.doe@example.com\r\n" +
		"To: roger.roe@example.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")
	auth := smtp.PlainAuth("", user, pass, host)

	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Printf("Error sending mail: %s", err)
	}

}
