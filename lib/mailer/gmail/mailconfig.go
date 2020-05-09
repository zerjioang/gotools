// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package gmail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"

	"github.com/zerjioang/gotools/lib/logger"
)

type MailServerConfig struct {
	Username    string
	Password    string
	EmailServer string
	Auth        smtp.Auth
}

func (c *MailServerConfig) SendWithSSL(fromMail, toMail *mail.Address, subject, body string) error {

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = subject

	headers["MIME-Version"] = "1.0"
	headers["X-Send-Timestamp"] = fmt.Sprintf(time.Now().Format(time.RFC850))
	headers["Content-Type"] = "text/html;charset=UTF-8"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.gmail.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("",
		c.Username,
		c.Password,
		host,
	)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, dialErr := tls.Dial("tcp", servername, tlsconfig)
	if dialErr != nil {
		logger.Debug("failed to dial tcp email server", dialErr)
		return dialErr
	}

	emailClient, clientErr := smtp.NewClient(conn, host)
	if clientErr != nil {
		logger.Error("failed to create a client connection", clientErr)
		return clientErr
	}

	// Auth
	if err := emailClient.Auth(auth); err != nil {
		logger.Error("failed to authenticate against email server", err)
		return err
	}

	// To && From
	if err := emailClient.Mail(fromMail.Address); err != nil {
		logger.Debug("failed to send email", err)
		return err
	}

	if err := emailClient.Rcpt(toMail.Address); err != nil {
		logger.Error("failed to rcpt email", err)
		return err
	}

	// Data
	w, err := emailClient.Data()
	if err != nil {
		logger.Error("failed to get email data", err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		logger.Error("failed to write email message", err)
		return err
	}

	err = w.Close()
	if err != nil {
		logger.Error("failed to close email connection", err)
		return err
	}

	qErr := emailClient.Quit()
	if qErr != nil {
		logger.Error("failed to quit email client", qErr)
	}
	return qErr
}

func (c *MailServerConfig) SendInsecure(fromMail, toMail *mail.Address, subject, msg string) error {

	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = subject

	headers["MIME-Version"] = "1.0"
	headers["X-Send-Timestamp"] = fmt.Sprintf(time.Now().Format(time.RFC850))
	headers["Content-Type"] = "text/html;charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + msg //+ base64.StdEncoding.EncodeToString([]byte(msg))

	err := smtp.SendMail(c.EmailServer, c.Auth, fromMail.Address, []string{toMail.Address}, []byte(message))
	if err != nil {
		logger.Error("An error occurred while sending email: %s" + err.Error())
	}
	return err
}

func encodeRFC2047(title string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Name: title}
	encoded := strings.Trim(addr.String(), " <@>")
	return encoded
}
