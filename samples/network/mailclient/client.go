package mailclient

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

type MailClient struct {
	ServerName string
	SmtpAddr string
	UserName string
	Password string
}

// Create new mail client. If your account use 2-fa authentication, you need to create password in for this device.
// ex for google mail server: https://stackoverflow.com/questions/60701936/error-invalid-login-application-specific-password-required/60718806#60718806
func NewMailClient(serverName, smtpAddr, username, password string) *MailClient {
	newClient := &MailClient{
		ServerName: serverName,
		SmtpAddr: smtpAddr,
		UserName: username,
		Password: password,
	}
	
	return newClient
}

func (m *MailClient) SendMail(subject, message, to string, cc, bcc[]string) (err error) {
	logger := rolling.New()

	// Set up the SMTP server connection
	auth := smtp.PlainAuth("", m.UserName, m.Password, m.ServerName)
    // Authenticate and send the email
	client, err := smtp.Dial(m.SmtpAddr)
	if err != nil {
		logger.Error("smtp.Dial failed", zap.Error(err))
		return
	}
	defer client.Quit()

    tlsConfig := &tls.Config{
		ServerName: m.ServerName,
	} 
	if err = client.StartTLS(tlsConfig); err != nil {
		logger.Error("client.StartTLS failed", zap.Error(err))
		return
	}

	if err = client.Auth(auth); err != nil {
		logger.Error("client.Auth failed", zap.Error(err))
		return
	}

	if err = client.Mail(m.UserName); err != nil {
		logger.Error("client.Mail failed", zap.Error(err))
		return
	}

	if err = client.Rcpt(to); err != nil {
		logger.Error("client.Rcpt failed", zap.Error(err))
		return
	}

	for _, ccAddress := range cc {
		if err = client.Rcpt(ccAddress); err != nil {
			logger.Error("client.Rcpt loop failed", zap.Error(err))
			return
		}
	}
	
	writer, err := client.Data(); if err != nil {
		fmt.Println(err)
		return
	}

	data := m.composeMail(to, subject, message, cc, bcc)
	if _, err = writer.Write(data); err != nil {
		fmt.Println(err)
		return
	}

	logger.Info("mail sent", zap.String("from", m.UserName), zap.String("to", to), zap.Strings("cc", cc))
	return nil
}

func (m *MailClient) composeMail(to, subject, message string, cc, bcc []string) []byte {
	haveCc := false
	haveBcc := false
	if len(cc) != 0 {
		haveCc = true	
	}
	if len(bcc) != 0 {
		haveBcc = true	
	}
	if haveBcc && haveCc {
		return []byte("To: " + to + "\r\n" +
			"Cc: " + m.commaSeparatedList(cc) + "\r\n" +
			"Bcc: " + m.commaSeparatedList(bcc) + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			message + "\r\n")

	} else if haveCc {
		return []byte("To: " + to + "\r\n" +
			"Cc: " + m.commaSeparatedList(cc) + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			message + "\r\n")
	} else if haveBcc {
		return []byte("To: " + to + "\r\n" +
			"Bcc: " + m.commaSeparatedList(bcc) + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			message + "\r\n")
	} else {
		return []byte("To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			message + "\r\n")
	}
}

func (m *MailClient) commaSeparatedList(list []string) string {
	return fmt.Sprintf("%s", list)
}