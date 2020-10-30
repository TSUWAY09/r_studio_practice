
package communication

import (
	"fmt"
	"log"
	"net/smtp"
	"stockScanner/types"
)

/*
	SnedMail - sends mail to the registered email ids
*/
func SendMail(c *types.Config) error {

	from := c.Mail.From
	password := c.Mail.Pass

	// toList is list of email address that email is to be sent.
	toList := c.Mail.To

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail smtp.gmail.com.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := c.Mail.Host

	// 587 is the default port of smtp server
	port := c.Mail.Port

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail