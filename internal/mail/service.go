package mail

import (
	"log"
	"time"

	"gopkg.in/mail.v2"
)

type MailConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    From     string
}

type MailService struct {
    dialer *mail.Dialer
    from   string
}

func NewMailService(cfg MailConfig) *MailService {
    d := mail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
		d.Timeout = 10 * time.Second // avoid hanging forever
    return &MailService{
        dialer: d,
        from:   cfg.From,
    }
}

func (m *MailService) Send(to, subject, body string) error {
    msg := mail.NewMessage()
    msg.SetHeader("From", m.from)
    msg.SetHeader("To", to)
    msg.SetHeader("Subject", subject)
    msg.SetBody("text/plain", body)

    if err := m.dialer.DialAndSend(msg); err != nil {
        log.Printf("Failed to send email: %v", err)
        return err
    }
    return nil
}
