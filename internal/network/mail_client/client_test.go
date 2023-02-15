package mailclient

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"
)

func TestSend(t *testing.T) {
    // Set the email recipient, sender, subject, and message
	to := "recipient@example.com"
	cc := []string{"cc1@example.com", "cc2@example.com"}
	from := "sender@gmail.com"
	password := "yourpassword"
	subject := "Test email"
	message := "This is a test email message."

	// Set up the SMTP server connection
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	smtpAddr := "smtp.gmail.com:587"
    // Authenticate and send the email
	client, err := smtp.Dial(smtpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Quit()

    tlsConfig := &tls.Config{
		ServerName: "smtp.gmail.com",
	} 
	if err = client.StartTLS(tlsConfig); err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Auth(auth); err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Mail(from); err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Rcpt(to); err != nil {
		fmt.Println(err)
		return
	}

	for _, ccAddress := range cc {
		if err = client.Rcpt(ccAddress); err != nil {
			fmt.Println(err)
			return
		}
	}

	data := []byte("To: " + to + "\r\n" +
		"Cc: " + commaSeparatedList(cc) + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")
	writer, err := client.Data(); if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = writer.Write(data); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent!")
}

func commaSeparatedList(list []string) string {
	return fmt.Sprintf("%s", list)
}
