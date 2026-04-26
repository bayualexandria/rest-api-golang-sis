package config

import (
    "os"
    "strconv"
)

type MailConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    From     string
}

func GetMailConfig() MailConfig {
    port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

    return MailConfig{
        Host:     os.Getenv("MAIL_HOST"),
        Port:     port,
        Username: os.Getenv("MAIL_USERNAME"),
        Password: os.Getenv("MAIL_PASSWORD"),
        From:     os.Getenv("MAIL_FROM"),
        
    }
}