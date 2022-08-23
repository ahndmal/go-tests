package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"testing"
)

func TestMail(t *testing.T) {
	// Connect to the remote SMTP server.
	//client, err := smtp.Dial("smtp.porkbun.com:587")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Set the sender and recipient first
	//if err := client.Mail("support@mail.com"); err != nil {
	//	log.Fatal(err)
	//}
	//if err := client.Rcpt("mail@gmail.com"); err != nil {
	//	log.Fatal(err)
	//}

	client, err := smtp.Dial("smtp.mailtrap.io:2525")
	if err != nil {
		log.Fatalf("Error when dialing connection %e", err)
	}

	//auth := smtp.PlainAuth("", os.Getenv("PRK_EMAIL"), os.Getenv("PRK_PASS"), "smtp.porkbun.com")
	auth := smtp.PlainAuth("", os.Getenv("MTRAP_USER"), os.Getenv("MTRAP_PASS"), "smtp.mailtrap.io")
	err2 := client.Auth(auth)
	if err2 != nil {
		log.Fatal(err2)
	}
	//err3 := smtp.SendMail("smtp://smtp.porkbun.com:465", auth, "support@bh.com", []string{"quadr@gmail.com"}, []byte("Hello!"))
	err3 := smtp.SendMail("smtp://smtp.mailtrap.io:2525", auth, "support@beastiehut.com", []string{"quadr988@gmail.com"}, []byte("Hello!"))
	if err3 != nil {
		log.Fatal(err2)
	}

	// Send the email body.
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = client.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
