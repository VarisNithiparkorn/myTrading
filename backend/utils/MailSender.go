package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"text/template"

	dto "github.com/VarisNithiparkorn/cryptoGraph/DTO"
	"github.com/VarisNithiparkorn/cryptoGraph/config"
)

var	(env = config.LoadEnv()
	username = env.MailUsername
	password = env.MailPassword
	smtpHost = env.SmtpHost
	smtpPort = env.SmtpPort
)
func SendTextMail(to []string, msg string) {
	
	auth := smtp.PlainAuth("",username,password,smtpHost)
	smtp.SendMail(smtpHost+":"+smtpPort,auth,username,to,[]byte(msg))
}

func SendHtmlMail(to []string, subject string, templatePath string, data dto.Response) error {

	var body bytes.Buffer
	

	tmpl, err := template.ParseFiles(templatePath) 
	if err != nil {
		log.Printf("Error parsing template file %s: %v", templatePath, err)
		return fmt.Errorf("failed to parse template: %w", err)
	}

	err = tmpl.Execute(&body, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return fmt.Errorf("failed to execute template: %w", err)
	}

	toHeader := strings.Join(to, ", ")
	auth := smtp.PlainAuth("",username,password,smtpHost)
	mimeHeaders := "MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n"

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"%s"+ 
		"\r\n"+
		"%s", 
		username,
		toHeader,
		subject,
		mimeHeaders,
		body.String()) 
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, username, to, []byte(msg))
}