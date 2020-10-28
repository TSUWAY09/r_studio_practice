
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
