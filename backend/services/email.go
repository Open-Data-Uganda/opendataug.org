package services

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

type Info struct {
	Email       string
	Token       string
	MailType    string
	UserName    string
	CurrentYear int
	Type        string
}

func (info Info) SendEmail() error {

	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

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
		return err
	}

	buff := new(bytes.Buffer)
	if err := t.Execute(buff, info); err != nil {
		log.Println(err)
		return err
	}

	result := buff.String()

	params := &resend.SendEmailRequest{
		To:      []string{info.Email},
		From:    os.Getenv("FROM_EMAIL"),
		Html:    result,
		Subject: info.MailType,
	}

	_, err = client.Emails.Send(params)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}
