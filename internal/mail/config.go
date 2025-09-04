package mail

import (
	"os"
	"strconv"
)

func LoadConfig() MailConfig {
    port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
    return MailConfig{
        Host:     os.Getenv("SMTP_HOST"),
        Port:     port,
        Username: os.Getenv("SMTP_USER"),
        Password: os.Getenv("SMTP_PASS"),
        From:     os.Getenv("SMTP_FROM"),
    }
}
