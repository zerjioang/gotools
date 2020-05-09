// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package model

import "encoding/json"

// current supported methods are:
// * SendGridMailDelivery
// * GmailDelivery
type EmailSenderMechanism func(maildata *Maildata) (json.RawMessage, error)

type MailAddress struct {
	User  string
	Email string
}

type Maildata struct {
	From      MailAddress
	To        MailAddress
	Cc        []MailAddress
	Bcc       []MailAddress
	Subject   string
	Plaintext string
	Htmltext  string
}
