package net

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"testing"
)

func TestMail(t *testing.T) {
	/*
		SMTP	smtp.porkbun.com	587	STARTTLS
		SMTP	smtp.porkbun.com	465	Implicit TLS
		IMAP	imap.porkbun.com	993	SSL (SSL/TLS)
		POP	pop.porkbun.com	995	SSL (SSL/TLS)
	*/

	// Connect to the remote SMTP server.
	//client, err := smtp.Dial("smtp.porkbun.com:587")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// Set the sender and recipient first
	//if err := client.Mail("support@mail.com"); err != nil {
	//	log.Fatal(err)
	//}
	//if err := client.Rcpt("mail@gmail.com"); err != nil {
	//	log.Fatal(err)
	//}
	smtpHost := "smtp.porkbun.com" //:465
	//imapHost := "imap.porkbun.com:993"
	client, err := smtp.Dial(smtpHost + ":465")
	if err != nil {
		log.Fatalf("Error when dialing connection %e", err)
	}

	auth := smtp.PlainAuth("", os.Getenv("PRK_EMAIL"), os.Getenv("PRK_PASS"), smtpHost)
	//auth := smtp.PlainAuth("", os.Getenv("MTRAP_USER"), os.Getenv("MTRAP_PASS"), "smtp.mailtrap.io")
	err2 := client.Auth(auth)
	if err2 != nil {
		log.Fatal(err2)
	}
	//err3 := smtp.SendMail("smtp://smtp.porkbun.com:465", auth, "support@bh.com", []string{"quadr@gmail.com"}, []byte("Hello!"))
	err3 := smtp.SendMail(fmt.Sprintf("%s:465", smtpHost), auth, os.Getenv("PRK_EMAIL"), []string{"quadr988@gmail.com"}, []byte("Hello!"))
	if err3 != nil {
		log.Fatal(err3)
	}

	// Send the email body.
	wc, err4 := client.Data()
	if err != nil {
		log.Fatal(err4)
	}
	_, err5 := fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err5)
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
