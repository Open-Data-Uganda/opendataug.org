package services

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Info struct {
	Email       string
	Token       string
	MailType    string
	UserName    string
	CurrentYear int
	Type        string
}

type SMTPCredentials struct {
	Email    string
	Password string
	Port     int
	Host     string
	Username string
}

func (info Info) SendEmail() error {
	templateUrl := ""
	if info.Type == "registration" {
		templateUrl = "./templates/registration.html"
	} else if info.Type == "invitation" {
		templateUrl = "./templates/invitation.html"
	} else if info.Type == "password_reset" {
		templateUrl = "./templates/reset_password.html"
	} else if info.Type == "quotation" {
		templateUrl = "./templates/quotation.html"
	} else {
		templateUrl = "./templates/reset_password.html"
	}

	t, err := template.ParseFiles(templateUrl)
	if err != nil {
		log.Println(err)
	}

	buff := new(bytes.Buffer)
	if err := t.Execute(buff, info); err != nil {
		log.Println(err)
	}

	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, _ := strconv.Atoi(smtpPortStr)

	smtpCreds := SMTPCredentials{
		Email:    os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Port:     smtpPort,
		Host:     os.Getenv("SMTP_HOST"),
		Username: os.Getenv("SMTP_USERNAME"),
	}

	result := buff.String()

	m := gomail.NewMessage()
	m.SetHeader("From", smtpCreds.Email)
	m.SetHeader("To", info.Email)
	m.SetHeader("Subject", info.MailType)
	m.SetBody("text/html", result)

	d := gomail.NewDialer(smtpCreds.Host, smtpCreds.Port, smtpCreds.Username, smtpCreds.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
	}

	return nil
}
