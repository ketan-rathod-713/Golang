package mail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func (s *service) SendMail(receiverEmailId string, htmlContent string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.SenderEmailId)
	m.SetHeader("To", receiverEmailId)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/html", htmlContent)

	d := gomail.NewDialer("smtp.gmail.com", 587, s.SenderEmailId, s.SMTP_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
