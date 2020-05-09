// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package sendgrid

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/zerjioang/gotools/lib/httpclient"
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/lib/mailer/model"
	"github.com/zerjioang/gotools/util/str"
)

/*
Sendgrid API email sender
More info at: https://sendgrid.com/docs/API_Reference/Web_API_v3/Mail/index.html

curl --request POST \
--url https://api.sendgrid.com/v3/mail/send \
--header "Authorization: Bearer $SENDGRID_API_KEY" \
--header 'Content-Type: application/json' \
--data '{
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
      "value": "Hello, World!"
    }
  ],
  "categories": ["noreply", "auth"]
}
'

*/

const (
	sendgridUrl = "https://api.sendgrid.com/v3/mail/send"
	payloadBody = `{
  "personalizations": [
    {
      "to": [
        {
          "name": "{{user}}",
          "email": "{{email}}"
        }
      ],
      "subject": "{{subject}}"
    }
  ],
  "from": {
    "name": "Etherniti",
    "email": "noreply@domain.tld"
  },
  "content": [
    {
      "type": "text/plain",
      "value": "{{plaintext}}"
    },
    {
      "type": "text/html",
      "value": "{{htmltext}}"
    }
  ],
  "categories": ["noreply", "auth"]
}`
)

var (
	apiKey               = ""
	defaultRequestHeader http.Header
	noApiKeyErr          = errors.New("no SENDGRID_API_KEY was defined")
	spaceRemover         = regexp.MustCompile(`\s+`)
)

func init() {
	//generate header used in all sendgrid request
	defaultRequestHeader = http.Header{
		"Content-Type":  []string{httpclient.ApplicationJSON},
		"Authorization": []string{"Bearer " + apiKey},
	}
	logger.Debug("creating sendgrid api client")
}

// todo make thread safe
func SetKey(k string) {
	logger.Debug("setting SENDGRID_API_KEY")
	apiKey = k
}

func buildSendGridPayload(maildata *model.Maildata) string {
	logger.Debug("building email payload data in sendgrid format")
	current := payloadBody
	current = strings.Replace(current, "{{user}}", maildata.To.User, 1)
	current = strings.Replace(current, "{{email}}", maildata.To.Email, 1)
	current = strings.Replace(current, "{{subject}}", maildata.Subject, 1)
	current = strings.Replace(current, "{{plaintext}}", maildata.Plaintext, 1)
	// we need to clean html tex. remove new lines
	maildata.Htmltext = cleanHtml(maildata.Htmltext)
	current = strings.Replace(current, "{{htmltext}}", maildata.Htmltext, 1)
	return current
}
func SendGridMailDelivery(data *model.Maildata) (json.RawMessage, error) {
	if apiKey == "" {
		logger.Error("aborting email delivery because no api key was defined in current environment variables")
		return nil, noApiKeyErr
	} else {
		emailStr := buildSendGridPayload(data)
		logger.Debug("sending email via sendgrid api")
		return httpclient.MakePost(nil, sendgridUrl, defaultRequestHeader, str.UnsafeBytes(emailStr))
	}
}

func cleanHtml(html string) string {
	html = strings.Replace(html, "\t", "", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "\n", "", -1)
	html = spaceRemover.ReplaceAllString(html, " ")
	html = strings.Replace(html, "> <", "><", -1)
	html = strings.Replace(html, "\"", "'", -1)
	return html
}
