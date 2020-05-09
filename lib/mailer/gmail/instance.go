package gmail

import (
	"net/smtp"
	"sync"
)

var (
	packageInitInstance       *MailServerConfig
	packageInstance           *MailServerConfig
	packageInstanceThreadSafe *MailServerConfig
	once                      sync.Once
)

/*
thread safe initialization. it only writes content to variable packageInitInstance once and
since this variable cannot be accesses by other code outside this file, it is completely thread safe
*/
func Init(user, password, server string) {
	packageInitInstance = newInternalServerConfig(user, password, server)
}

func newInternalServerConfig(user, password, server string) *MailServerConfig {
	conf := new(MailServerConfig)
	conf.Username = user
	conf.Password = password
	conf.EmailServer = server
	conf.Auth = smtp.PlainAuth("",
		conf.Username,
		conf.Password,
		server, //server should only include server name
	)
	return conf
}

func GetGmailServerConfigInstanceInit() *MailServerConfig {
	return packageInitInstance
}

/*
@deprecated use GetGmailServerConfigInstanceInit() instead

warning: this is a non-thread safe implementation of singleton.
use at your own risk knowning variable ioproto access
*/
func GetGmailServerConfigInstance(user, password, server string) *MailServerConfig {
	//not thread safe code
	if packageInstance == nil {
		packageInstance = newInternalServerConfig(user, password, server)
	}
	return packageInstance
}

/*
@deprecated use GetGmailServerConfigInstanceInit() instead
*/
func GetGmailServerConfigInstanceThreadSafe(user, password, server string) *MailServerConfig {
	//thread safe code
	once.Do(func() {
		packageInstanceThreadSafe = newInternalServerConfig(user, password, server)
	})
	return packageInstanceThreadSafe
}
