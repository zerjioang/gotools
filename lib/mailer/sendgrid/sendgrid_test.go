// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package sendgrid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/mailer/model"
)

const (
	testdata = `{
  "personalizations": [
    {
      "to": [
        {
          "name": "User 01",
          "email": "user01@domain.tld"
        }
      ],
      "subject": "Hello, World!"
    }
  ],
  "from": {
    "name": "Etherniti",
    "email": "noreply@domain.tld"
  },
  "content": [
    {
      "type": "text/plain",
      "value": "Hello, World!"
    },
    {
      "type": "text/html",
      "value": "<h1>Hello, World!</h1"
    }
  ],
  "categories": ["noreply", "auth"]
}`
)

func TestSendGridSendEmail(t *testing.T) {
	t.Run("send-email", func(t *testing.T) {
		data := &model.Maildata{
			From:      model.MailAddress{User: "Etherniti", Email: "noreply@domain.tld"},
			To:        model.MailAddress{User: "User 02", Email: "user02@domain.tld"},
			Subject:   "SendGrid Mail Delivery Test",
			Plaintext: "this is a SendGrid Mail Delivery Test",
			Htmltext:  "<h1>this is a SendGrid Mail Delivery Test</h1>",
		}
		response, err := SendGridMailDelivery(data)
		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
