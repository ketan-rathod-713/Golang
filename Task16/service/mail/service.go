package mail

// all jwt auth services define here
type service struct {
	SenderEmailId string
	SMTP_PASSWORD string
}

func New(smtpPassword string, senderEmailId string) Service {
	return &service{
		SMTP_PASSWORD: smtpPassword,
		SenderEmailId: senderEmailId,
	}
}

type Service interface {
	SendMail(receiverEmailId string, htmlContent string) error
}
