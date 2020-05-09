// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package gmail

import (
	"encoding/json"
	"net/mail"

	"github.com/zerjioang/gotools/lib/mailer/model"
)

var (
	defaultGmailer *ApiGmailer
)

type ApiGmailer struct {
	maxLimit     int
	currentLimit int
	mailengine   *MailServerConfig
	from         *mail.Address
}

func NewApiGmailer() *ApiGmailer {
	em := new(ApiGmailer)
	// currently sent email amount
	em.currentLimit = 0
	//by default, google free email services only allow to send a maximum of 100 emails per day
	em.maxLimit = 100
	//default email sending address
	em.from = &mail.Address{Name: "test", Address: "noreply@test.org"}
	//initialize default mail engine
	em.mailengine = GetGmailServerConfigInstanceInit()
	return em
}

func init() {
	defaultGmailer = NewApiGmailer()
}

func SendWithGmail(data *model.Maildata) (json.RawMessage, error) {
	//send email
	err := defaultGmailer.mailengine.SendWithSSL(
		defaultGmailer.from,
		&mail.Address{Name: data.To.User, Address: data.To.Email},
		data.Subject,
		data.Htmltext,
	)
	return nil, err
}
