package fakesender

import (
	"encoding/json"
	"errors"

	"github.com/zerjioang/gotools/lib/uuid/randomuuid"

	"github.com/zerjioang/gotools/lib/fs"
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/lib/mailer/model"
)

func FakeEmailSender(maildata *model.Maildata) (json.RawMessage, error) {
	if maildata == nil {
		logger.Debug("sending email...")
		return nil, errors.New("no email data")
	}
	filename := randomuuid.GenerateIDString()
	logger.Debug("creating temp file: ", filename)
	return nil, fs.WriteFile("/tmp/"+filename.String()+".html", maildata.Htmltext)
}
